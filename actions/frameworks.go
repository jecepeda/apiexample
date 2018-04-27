package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/jcepedavillamayor/apiexample/models"
	"github.com/pkg/errors"
)

type FrameworksResource struct {
	buffalo.Resource
}

func (v FrameworksResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	frameworks := &models.Frameworks{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Frameworks from the DB
	if err := q.All(frameworks); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.JSON(frameworks))
}

// Show gets the data for one Framework. This function is mapped to
// the path GET /frameworks/{framework_id}
func (v FrameworksResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Framework
	framework := &models.Framework{}

	// To find the Framework the parameter framework_id is used.
	if err := tx.Find(framework, c.Param("framework_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.JSON(framework))
}

// Create adds a Framework to the DB. This function is mapped to the
// path POST /frameworks
func (v FrameworksResource) Create(c buffalo.Context) error {
	// Allocate an empty Framework
	framework := &models.Framework{}

	// Bind framework to the html form elements
	if err := c.Bind(framework); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(framework)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		return c.Error(422, verrs)
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Framework was created successfully")

	// and redirect to the frameworks index page
	return c.Render(201, r.JSON(framework))
}

// Update changes a Framework in the DB. This function is mapped to
// the path PUT /frameworks/{framework_id}
func (v FrameworksResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Framework
	framework := &models.Framework{}

	if err := tx.Find(framework, c.Param("framework_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Framework to the html form elements
	if err := c.Bind(framework); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(framework)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.JSON(verrs))
	}

	// and redirect to the frameworks index page
	return c.Render(200, r.JSON(framework))
}

// Destroy deletes a Framework from the DB. This function is mapped
// to the path DELETE /frameworks/{framework_id}
func (v FrameworksResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Framework
	framework := &models.Framework{}

	// To find the Framework the parameter framework_id is used.
	if err := tx.Find(framework, c.Param("framework_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(framework); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Framework was destroyed successfully")

	// Redirect to the frameworks index page
	return c.Render(200, r.JSON(framework))
}
