package controller

// Controller holds the controllers for the entire across application
type Controller struct {
	Customer interface{ Customer }
	Product  interface{ Product }
	Order    interface{ Order }
}
