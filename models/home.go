package models

import "encoding/json"

type Education struct {
	ID          int    `json:"id"`
	Year        string `json:"year"`
	DegName     string `json:"degname"`
	CollegeName string `json:"collegename"`
	Grade       string `json:"grade"`
}

// For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var EducationList = []Education{
	{ID: 1, Year: "2018-2020", DegName: "MSc IT", CollegeName: "Gonsalo Garcia College, Vasai", Grade: ""},
	{ID: 2, Year: "2015-2018", DegName: "BSc IT", CollegeName: "Wilson College, Mumbai", Grade: "CGPA: 6.67 / 7"},
	{ID: 3, Year: "2013-2015", DegName: "HSC", CollegeName: "Thomas Baptista J C, Vasai", Grade: "Percentege: 75.85"},
	{ID: 4, Year: "2012-2013", DegName: "SSC", CollegeName: "St. Anthonys Convent High School, Vasai", Grade: "Percentege: 82.36"},
}

// Return a list of all the articles
func GetEducationList() []Education {
	return EducationList
}

type Experience struct {
	ID      int    `json:"id"`
	Year    string `json:"year"`
	JobName string `json:"jobname"`
	OrgName string `json:"orgname"`
	Desc    string `json:"desc"`
}

// For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var ExperienceList = []Experience{
	{
		ID:      1,
		Year:    "JULY 2019-PRESENT",
		JobName: "Sr. Software Engineer",
		OrgName: "CHRISEL TECHNOLAB",
		Desc:    "Python Developer working with libraries like Flask, Flask-restful, Pandas, Numpy, Pyqt5, Merklelib, Sockets and many more to fulfill the requirements of the company. Also having a decent exposure with Blockchain and it’s underlying concepts.",
	},
	{
		ID:      2,
		Year:    "JULY 2018-NOV 2018",
		JobName: "Jr. Software Engineer",
		OrgName: "DQUIP",
		Desc:    "Backend web developer working with JavaScript/jQuery, PHP, Laravel, MYSQL. Responsibilities include developing solutions for CRM Web Applications for different industries.",
	},
}

// Return a list of all the experience
func GetExperienceList() []Experience {
	return ExperienceList
}

var ProfessionList = []string{"Python Developer.", "Go Developer.", "Backend Developer.", "Gamer."}

// Return a list of all the professions
func GetProfessionList() string {
	jsonStr, err := json.MarshalIndent(ProfessionList, "", "  ")
	if err != nil {
		// Handle the error
	}
	return string(jsonStr)
}

type Project struct {
	ID          int    `json:"id"`
	Year        string `json:"year"`
	ProjectName string `json:"projectname"`
	Language    string `json:"language"`
	Desc        string `json:"desc"`
}

var ProjectList = []Project{
	{
		ID:          1,
		Year:        "JULY 2019-PRESENT",
		ProjectName: "go",
		Language:    "proj 1",
		Desc:        "blablabla.",
	},
	{
		ID:          2,
		Year:        "JULY 2018-NOV 2018",
		ProjectName: "Proj 2",
		Language:    "python",
		Desc:        "babababa.",
	},
}

// Return a list of all the experience
func GetProjectList() []Project {
	return ProjectList
}

type MyInfo struct {
	BrandFirst        string `json:"brandfirst"`
	BrandLater        string `json:"brandlater"`
	FirstName         string `json:"firstname"`
	LastName          string `json:"lastname"`
	AboutMe           string `json:"aboutme"`
	DOB               string `json:"dob"`
	Address           string `json:"address"`
	FullAddress       string `json:"fulladdress"`
	ZipCode           string `json:"zipcode"`
	PersonalEmail     string `json:"personalemail"`
	PersonalMobile    string `json:"personalmobile"`
	Website           string `json:"website"`
	ProjectsCompleted string `json:"projectscompleted"`
}

var MyInfoData = MyInfo{
	BrandFirst:        "E",
	BrandLater:        "rson",
	FirstName:         "Erson",
	LastName:          "Pereira",
	AboutMe:           "Hi hellow loreum ipsium",
	DOB:               "March 21, 1998",
	Address:           "Vasai MH India",
	FullAddress:       "New Chawri Wadi, Saloli, Vasai West",
	ZipCode:           "401201",
	PersonalEmail:     "ersonpereiracr7@gmail.com",
	PersonalMobile:    "+91 8999696267",
	Website:           "www.ersonpereira.com",
	ProjectsCompleted: "2",
}

// Return a list of all the experience
func GetMyInfoData() MyInfo {
	return MyInfoData
}
