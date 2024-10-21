package main

import "Goweb/frame"

func SubjectAddController(c *frame.Context) error {
	c.SetOkStatus().Json("ok, SubjectAddController")
	return nil
}

func SubjectListController(c *frame.Context) error {
	c.SetOkStatus().Json("ok, SubjectListController")
	return nil
}

func SubjectDelController(c *frame.Context) error {
	c.SetOkStatus().Json("ok, SubjectDelController")
	return nil
}

func SubjectUpdateController(c *frame.Context) error {
	c.SetOkStatus().Json("ok, SubjectUpdateController")
	return nil
}

func SubjectGetController(c *frame.Context) error {
	c.SetOkStatus().Json("ok, SubjectGetController")
	return nil
}

func SubjectNameController(c *frame.Context) error {
	c.SetOkStatus().Json("ok, SubjectNameController")
	return nil
}
