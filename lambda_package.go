package kanarya

type LambdaBucket struct {
	Name string
	Key  string
}

type LambdaFunction struct {
	Name string
}

type LambdaAlias struct {
	Name string
}

type LambdaPackage struct {
	Location string
	Function LambdaFunction
	Bucket   LambdaBucket
	Alias    LambdaAlias
}
