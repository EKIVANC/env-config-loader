# What is env-config-loader ?
`env-config-loader` library is a simple and plain way of managing environment variables.

## Install

```shell
go get github.com/ekivanc/env-config-loader
```


## How to use the library:

* Write your environment variables in a config file with the following format:

```
MyKeyName=MyCoolValue
OtherKeyName=OtherCoolValue
```

* each line in config file will be considered as a different environment parameter.

* if you want to add some description for your environment variable, you can write commnets by begining `#` character:

```yaml
# here is a sample comment line
SampleKey=SuperCoolValue
```

* To load the environment variables you can call `LoadEnvVariables` function:

```go
LoadEnvVariables("./configs/app.config")
```

* you can use golang standart way for reading your values:
```
os.Getenv("MyKeyName")
```


## Why this library needed ?
#### The common issues with other popular env parameter management libs are:

* Golang already supports getting and setting environment variables natively, in general what you only need is keeping all those variables in a single config file.
Other popular libraries also use some other 3rd party libraries, which you really do not need (or want) in your application, it will make your deployable artifacts bigger, and more fragile because of dependency tree.  

* Other popular libraries also make some type conversions inside library, so you can not really %100 sure what you will have when you change the env variable value.

* Environent variables management should be simple, if you really need to change a value in application level, then what you need is probably not an environment variable, but a value needs to be stored in a datastore (like a database) or if you deploy your app in a cloud managed environment (like aws or gcp..) then you may consider using your cloud providers parameters store.



So if what you only need is to keep the required env variables in a config file and then loading them at runtime, then this lib is what you just need.
