package models

func (as *ModelSuite) Test_FrameworkSuccess() {
	f := &Framework{
		Title:       "Foobar",
		Description: "FOOBARFOOBAR",
	}
	if _, err := as.DB.ValidateAndCreate(f); err != nil {
		as.Error(err)
	}
	actual := &Framework{}
	if err := as.DB.Find(actual, f.ID); err != nil {
		as.Error(err)
	}
	as.Equal(f.ID, actual.ID)
	as.Equal(f.Title, actual.Title)
	as.Equal(f.Description, actual.Description)
}

func (as *ModelSuite) Test_FrameworkErrorMissingFields() {
	f := &Framework{
		Title: "Foobar",
	}
	if verrs, err := as.DB.ValidateAndCreate(f); err != nil || verrs != nil {
		as.EqualError(verrs, "Description can not be blank.")
	} else {
		as.Fail("The framework can be created")
	}
	f = &Framework{
		Description: "FOOBARFOOBAR",
	}
	if verrs, err := as.DB.ValidateAndCreate(f); err != nil || verrs != nil {
		as.EqualError(verrs, "Title can not be blank.")
	}
}
