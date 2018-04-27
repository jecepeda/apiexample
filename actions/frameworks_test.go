package actions

import (
	"encoding/json"

	"github.com/jcepedavillamayor/apiexample/models"
)

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
	as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_FrameworksResource_Destroy() {
	as.Fail("Not Implemented!")
}
