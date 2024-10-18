package main

import "Goweb/frame"

func SubjectAddController(c *frame.Context) error {
	c.Json(200, "ok, SubjectAddController")
	return nil
}

func SubjectListController(c *frame.Context) error {
	c.Json(200, "ok, SubjectListController")
	return nil
}

func SubjectDelController(c *frame.Context) error {
	c.Json(200, "ok, SubjectDelController")
	return nil
}

func SubjectUpdateController(c *frame.Context) error {
	c.Json(200, "ok, SubjectUpdateController")
	return nil
}

func SubjectGetController(c *frame.Context) error {
	c.Json(200, "ok, SubjectGetController")
	return nil
}

func SubjectNameController(c *frame.Context) error {
	c.Json(200, "ok, SubjectNameController")
	return nil
}
