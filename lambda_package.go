package kanarya

// A LambdaBucket is used to reflect structure of an AWS S3 bucket that is used
// to store lambda functions.
type LambdaBucket struct {
	Name string
	Key  string
}

// A LambdaFunction is used to reflect structure of an AWS Lambda function.
type LambdaFunction struct {
	Name string
}

// A LambdaAlias is used to reflect structure of a lambda alias. Lambda aliases
// are used in traffic shifthing and gradual deployments.
type LambdaAlias struct {
	Name string
}

// A LambdaPackage is an internal representation of a lambda function that can
// be deployed gradually. LambdaPackage is a composite struct made of other
// structs.
type LambdaPackage struct {
	Location string
	Function LambdaFunction
	Bucket   LambdaBucket
	Alias    LambdaAlias
}
