# configd-injector

The configd-injector adds your config'd configuration to your project as a dependency.

# Downloads

- [Linux Binary](https://github.com/cheikhshift/configd-injector/raw/master/configd-injector.tar.gz)
	- md5 checksum of binary : c1b34eddb4aa769ca2b9c8c49c3252c7

- [Windows Binary](https://github.com/cheikhshift/configd-injector/raw/master/configd-injector.exe.zip)
	- md5 checksum of binary : 7541adc63cb1c8cdc3702d36a0b3a431

# Install binary on linux

### Download binary

Run the following command to download binary

	curl  https://github.com/cheikhshift/configd-injector/raw/master/configd-injector.tar.gz \
  	--output configd.tar.gz

### Decompress archive

	tar -pxvzf configd.tar.gz

### Install command

	sudo mv configd-injector /usr/sbin/


# Requirements (to build)

- Go 1.8+
- Environment variable Path has `$GOPATH/bin` in it.
- Config'd API key. Find it [here](https://configd.gophersauce.com/login)

# Get source

	go get github.com/cheikhshift/configd-injector


### Usage

Prior to running your application, run the following command to download your configuration and add it as a dependency.


#### NodeJS project

Prepare for Node :

	configd-injector -key=API_KEY -path=~/project -node

You can retrieve configuration data by using the following module : `require("configd")`. This module will be created within your project's package folder. The module is an Object with your configuration data.

#### Go Project

Prepare for Go :

	configd-injector -key=API_KEY -path=~/project -go

You can import your configuration data with the following package path : `import "configd"`.

To access data, configd package has an exported variable named Settings with your configuration data with type `map[string]interface{}`.

You must assert your types, Settings has type `map[string]interface{}`. The configd package features a set of exported helper functions to help with interface assertion. Here is a list of the configd package functions :

- func Int(i interface{}) int64 
- func Bool(i interface{}) bool 
- func String(i interface{}) string 
- func Map(i interface{}) map[string]interface{} 

 
 Find more integration guides at the [wiki](https://github.com/cheikhshift/configd/wiki)


## Roadmap

- [] Support PHP composer.
- [] Support Ruby Gems.
- [x] Support Java projects.
- [] Manage lifecycle of web application.
- [] Kubernetes operator, to enable live reloads.
- [] Docker container image.
  - To enable live reloads.
  - To feature deployment without rebuilding docker image.

### Contribution

PRs are encouraged and appreciated. 
