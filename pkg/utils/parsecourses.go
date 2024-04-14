package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type StepikCourseList_ struct {
	Courses []int `json:"courses"`
}

type StepikCourseLists_ struct {
	Meta        StepikMeta_         `json:"meta"`
	CourseLists []StepikCourseList_ `json:"course-lists"`
}

type StepikCourse_ struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Summary       string `json:"summary"`
	Cover         string `json:"cover"`
	ReviewSummary int    `json:"review_summary"`
	LearnersCount int    `json:"learners_count"`
}

type StepikCourseJson_ struct {
	Courses []StepikCourse_ `json:"courses"`
}

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

type StepikMeta_ struct {
	HasNext bool `json:"has_next"`
}

func getStepikCourses() []StepikCourse_ {
	page := 1
	var coursesIds []string
	var courses []StepikCourse_
	for {
		resp, err := http.Get("https://stepik.org/api/course-lists?page=" + strconv.Itoa(page))
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body)
		var cl StepikCourseLists_
		if json.Unmarshal([]byte(sb), &cl) != nil {
			log.Println("JSON decode error!")
			return []StepikCourse_{}
		}
		if cl.Meta.HasNext == false {
			break
		}
		for _, category := range cl.CourseLists {
			for _, cur_id := range category.Courses {
				strid := strconv.Itoa(cur_id)
				if !contains(coursesIds, strid) {
					coursesIds = append(coursesIds, strid)
				}
			}
		}
		if page > 1 {
			break
		}
		page += 1
	}

	for i, courseId := range coursesIds {
		resp, err := http.Get("https://stepik.org/api/courses?ids%5B%5D=" + courseId)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body)
		var cl StepikCourseJson_
		if json.Unmarshal([]byte(sb), &cl) != nil {
			log.Println("JSON decode error!")
			log.Println(json.Unmarshal([]byte(sb), &cl))
			return []StepikCourse_{}
		}
		courses = append(courses, cl.Courses[0])
		if i > 100 {
			break
		}
	}
	return courses
}
