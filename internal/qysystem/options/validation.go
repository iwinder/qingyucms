package options

func (o *Options) Validate() []error {
	var errs []error

	errs = append(errs, o.GenericServerRunOptions.Validate()...)
	errs = append(errs, o.InsecureServing.Validate()...)
	errs = append(errs, o.MySQLOptions.Validate()...)
	errs = append(errs, o.Log.Validate()...)
	return errs
}
