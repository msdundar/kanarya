# Kanarya

> Kanarya is _canary_ in Turkish.

`kanarya` is a Go module that takes care of canary deployments in AWS Lambda.

This module acts as a wrapper on top of AWS Go SDK and makes easier to
implement canary deployments in your lambda projects.

`kanarya` can be used locally in your Go projects, can be implemented as a CLI
tool, or can be used on CI, depending on your needs.

## Install

![](https://raw.githubusercontent.com/msdundar/kanarya/master/assets/logo_256.png)

```sh
go get github.com/msdundar/kanarya
```

> `kanarya` uses Go Modules to manage dependencies.

## Usage

### Credentials

`kanarya` relies on AWS Shared Configuration. You need to set your credentials
before using the package.

- Set `AWS_REGION` environment variable to the default region.
- Either set `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` - or define your
  credentials in `~/.aws/credentials` file.

For more details follow official [AWS Guidelines](https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/).

### Implementation

Create S3 and Lambda instances to use in AWS operations:

```golang
s3Client     = kanarya.S3Client("AWS_REGION_TO_DEPLOY")
lambdaClient = kanarya.LambdaClient("AWS_REGION_TO_DEPLOY")
```

Create a deployment package by using `kanarya.LambdaPackage` struct:

```golang
lambdaPackage := kanarya.LambdaPackage{
  Location: "file/path/for/lambda/package/index.zip",
  Function: kanarya.LambdaFunction{
    Name: "YOUR_LAMBDA_NAME",
  },
  Bucket: kanarya.LambdaBucket{
    Name: "YOUR_BUCKET_NAME",
    Key:  "upload/folder/in/s3bucket/index.zip",
  },
  Alias: kanarya.LambdaAlias{
    Name: "YOUR_LAMBDA_ALIAS", // alias used by clients
  },
}
```

Upload deployment package to S3:

```golang
_, err := kanarya.UploadToS3(s3Client, lambdaPackage)
```

Update function code located in `$LATEST`:

```golang
_, err := kanarya.UpdateFunctionCode(lambdaClient, lambdaPackage)
```

Publish a new version for shifting traffic later:

```golang
resp, err := kanarya.PublishNewVersion(lambdaClient, lambdaPackage)
newVersion := resp.Version
```

Create a JSON payload to use in health check requests:

```golang
request := yourRequestStruct{Something: "some value"}
payload, err := json.Marshal(request)
```

> Adjust this JSON according to the expected request by your lambda.

Start gradual deployment:

```golang
oldVersion, err := kanarya.GradualRollOut(
  lambdaClient,
  lambdaPackage,
  newVersion,
  0.1000000, // roll out rate in each step. 0.1 equals to 10%.
  10, // number of health checks you would like to run on each step.
  60, // number of seconds to sleep on each step.
  payload,
)
```

If there are errors during the gradual rollout, auto-rollback to the previous
healthy version:

```golang
if err != nil {
  kanarya.FullRollOut(lambdaClient, lambdaPackage, oldVersion)
  os.Exit(1)
}
```

If gradual rollout is successful, then fully roll out the new version:

```golang
_, err := kanarya.FullRollOut(lambdaClient, lambdaPackage, newVersion)
```

And that's it! You can combine the example above for cross-regional deployments,
by updating the `s3Client` or `lambdaClient` on the fly with a new region.

## Contibutions

1. Fork the repo
2. Clone the fork (`git clone git@github.com:YOUR_USERNAME/kanarya.git && cd kanarya`)
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Make changes and add them (`git add --all`)
5. Commit your changes (`git commit -m 'Add some feature'`)
6. Push to the branch (`git push origin my-new-feature`)
7. Create a pull request

## License

See [MIT-LICENSE](https://github.com/msdundar/kanarya/blob/master/MIT-LICENSE).
