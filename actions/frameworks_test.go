package actions

import (
	"encoding/json"

	"fmt"

	"github.com/brianvoe/gofakeit"
	"github.com/jcepedavillamayor/apiexample/models"
)

func (as *ActionSuite) createRandomFrameworks(n int) models.Frameworks {
	frameworks := make(models.Frameworks, 0)
	for i := 0; i < n; i += 1 {
		f := &models.Framework{}
		gofakeit.Struct(f)
		if err := as.DB.Create(f); err != nil {
			as.Error(err)
		}
		frameworks = append(frameworks, *f)
	}
	return frameworks
}

func (as *ActionSuite) checkFrameworks(expected, actual models.Frameworks) {
	for _, e := range expected {
		match := models.Framework{}
		for _, a := range actual {
			if e.ID == a.ID {
				match = a
			}
		}
		if match == (models.Framework{}) {
			msg := fmt.Sprintf("framework %v not in list", e)
			as.Fail(msg)
		}
		as.Equal(e.Title, match.Title)
		as.Equal(e.Description, match.Description)
	}
}

func (as *ActionSuite) Test_FrameworksResource_List() {
	expected := as.createRandomFrameworks(4)

	response := as.JSON("/frameworks").Get()
	as.Equal(200, response.Code)

	actual := make(models.Frameworks, 0)
	err := json.Unmarshal(response.Body.Bytes(), &actual)

	if err != nil {
		as.Error(err)
	}

	as.checkFrameworks(expected, actual)
}

func (as *ActionSuite) Test_FrameworksResource_ListEmpty() {
	res := as.JSON("/frameworks").Get()
	as.Equal(200, res.Code)
	as.Equal("[]\n", res.Body.String())
}

func (as *ActionSuite) Test_FrameworksResource_Show() {
	f := models.Framework{
		Description: "FOOBARFOOBAR",
		Title:       "Foobar",
	}
	err := as.DB.Create(&f)
	if err != nil {
		as.Error(err)
	}
	// So now we add the model
	// We make the request
	res := as.JSON("/frameworks/" + f.ID.String()).Get()
	as.Equal(200, res.Code)

	// we check that the information is correct
	content := models.Framework{}
	err = json.Unmarshal(res.Body.Bytes(), &content)
	as.Equal(f.Title, content.Title)
	as.Equal(f.Description, content.Description)
	as.Equal(f.ID, content.ID)
}

func (as *ActionSuite) Test_FrameworksResource_Create() {
	f := &models.Framework{
		Title:       "FooBar",
		Description: "FOOBARFOOBAR",
	}
	res := as.JSON("/frameworks").Post(f)
	as.Equal(201, res.Code)

	// Check if the framework is in the DB
	err := as.DB.First(f)
	as.NoError(err)
	as.NotZero(f.ID)
	as.NotNil(f.CreatedAt)
	// check that the data inserted is correct
	as.Equal("FooBar", f.Title)
	as.Equal("FOOBARFOOBAR", f.Description)

	// check if the response contains the structure
	var response models.Framework
	err = json.Unmarshal(res.Body.Bytes(), &response)
	if err != nil {
		as.Error(err)
	}

	as.Equal(f.Title, response.Title)
	as.Equal(f.ID, response.ID)
	as.Equal(f.Description, response.Description)
}

func (as *ActionSuite) Test_FrameworksResource_CreateFailsNoTitle() {
	f := &models.Framework{
		Description: "FOOBARFOOBAR",
	}
	res := as.JSON("/frameworks").Post(f)
	// Since we send a correct structure but  we do not sent
	// the complete entity to set, we receibe 422 Unprocessable Entity response
	as.Equal(422, res.Code)

	var rest []byte
	_, err := res.Result().Body.Read(rest)
	if err != nil {
		as.Fail(err.Error())
	}

	expected := ErrorStruct{
		Status: 422,
		Error:  "Title can not be blank.",
	}

	real := ErrorStruct{}
	err = json.Unmarshal(res.Body.Bytes(), &real)
	if err != nil {
		as.Error(err)
	}
	as.Equal(expected, real)
}

func (as *ActionSuite) Test_FrameworksResource_CreateFailsNoDescription() {
	f := &models.Framework{
		Title: "FooBar",
	}
	res := as.JSON("/frameworks").Post(f)
	// Since we send a correct structure but  we do not sent
	// the complete entity to set, we receibe 422 Unprocessable Entity response
	as.Equal(422, res.Code)

	var rest []byte
	_, err := res.Result().Body.Read(rest)
	if err != nil {
		as.Fail(err.Error())
	}

	expected := ErrorStruct{
		Status: 422,
		Error:  "Description can not be blank.",
	}

	real := ErrorStruct{}
	err = json.Unmarshal(res.Body.Bytes(), &real)
	if err != nil {
		as.Error(err)
	}
	as.Equal(expected, real)
}

func (as *ActionSuite) Test_FrameworksResource_Update() {
	f := &models.Framework{
		Title:       "Foobar",
		Description: "FOOBARFOOBAR",
	}
	if err := as.DB.Create(f); err != nil {
		as.Error(err)
	}
	f.Title = "Foo"
	f.Description = "FOOBAR"
	res := as.JSON("/frameworks/" + f.ID.String()).Put(f)

	as.Equal(200, res.Code)

	actual := models.Framework{}
	err := json.Unmarshal(res.Body.Bytes(), &actual)
	if err != nil {
		as.Error(err)
	}
	as.Equal(f.ID, actual.ID)
	as.Equal(f.Title, actual.Title)
	as.Equal(f.Description, actual.Description)
}

func (as *ActionSuite) Test_FrameworksResource_Destroy() {
	f := &models.Framework{
		Title:       "Foobar",
		Description: "FOOBARFOOBAR",
	}
	if err := as.DB.Create(f); err != nil {
		as.Error(err)
	}
	res := as.JSON("/frameworks/" + f.ID.String()).Delete()

	as.Equal(200, res.Code)
	//assert that is not there anymore

	res = as.JSON("/frameworks/" + f.ID.String()).Delete()
	as.Equal(404, res.Code)
}
