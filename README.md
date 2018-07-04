## Statup Example Plugin
Since Statup is built in Go Language we can use the [Go Plugin](https://golang.org/pkg/plugin/) feature to create dynamic plugins that run on load. Statup has an event anytime anything happens, you can create your own plugins and do any type of function. To implement your own ideas into Statup, use the plugin using the [statup/plugin](https://github.com/hunterlong/statup/blob/master/plugin/main.go) package.
```
go get github.com/hunterlong/statup/plugin
```

<p align="center">
<img width="95%" src="https://s3-us-west-2.amazonaws.com/gitimgs/statupplugin.png">
</p>

## Building Plugins
Plugins don't need a push request and they can be private! You'll need to compile your plugin to the Golang `.so` binary format. Once you've built your plugin, insert it into the `plugins` folder in your Statup directory and reboot the application. Clone the [Example Statup Plugin](https://github.com/hunterlong/statup_plugin) repo and try to build it yourself!

#### Build Requirements
- You must have `main.go`
- You must create the Plugin variable on `init()`

```bash
git clone https://github.com/hunterlong/statup_plugin
cd statup-plugin
go build -buildmode=plugin -o example.so
```
###### Insert `example.so` into the `plugins` directory and reload Statup

## Testing Statup Plugins
Statup includes a couple tools to help you on your Plugin journey, you can use `statup test plugins` command to test all plugins in your `/plugins` folder. This test will attempt to parse your plugin details, and then it will send events for your plugin to be fired. 
```
statup test plugins
```
<p align="center">
<img width="75%" src="https://s3-us-west-2.amazonaws.com/gitimgs/statupplugintester.png">
</p>

Your plugin should be able to parse and receive events before distributing it. The test tools creates a temporary database (SQLite) that your plugin can interact with. Statup uses [upper.io/db.v3](https://upper.io/db.v3) for database interactions. The database is passed to your plugin `OnLoad(db sqlbuilder.Database)`, so you can use the `db` variable passed here.

## Statup Plugin Interface
Please remember Golang plugin's are very new and Statup plugin package may change and 'could' brake your plugin. Checkout the [statup/plugin package](https://github.com/hunterlong/statup/blob/master/plugin/main.go) to see the most current interfaces. 
```go
type PluginActions interface {
	GetInfo() Info
	GetForm() string
	SetInfo(map[string]interface{}) Info
	Routes() []Routing
	OnSave(map[string]interface{})
	OnFailure(map[string]interface{})
	OnSuccess(map[string]interface{})
	OnSettingsSaved(map[string]interface{})
	OnNewUser(map[string]interface{})
	OnNewService(map[string]interface{})
	OnUpdatedService(map[string]interface{})
	OnDeletedService(map[string]interface{})
	OnInstall(map[string]interface{})
	OnUninstall(map[string]interface{})
	OnBeforeRequest(map[string]interface{})
	OnAfterRequest(map[string]interface{})
	OnShutdown()
	OnLoad(sqlbuilder.Database)
}
```

## Event Parameters
All event interfaces for the Statup Plugin will return a `map[string]interface{}` type, this is because the plugin package will most likely update and change in the future, but using this type will allow your plugin to continue even after updates.

## Example of an Event
Knowing what happens during an event is important for your plugin. For example, lets have an event that echo something when a service has a Failure status being issued. Checkout some example below to see how this golang plugin action works. 

```go
func (p pkg) OnSuccess(data map[string]interface{}) {
    fmt.Println("Statup Example Plugin received a successful service hit! ")
    fmt.Println("Name:    ", data["Name"])
    fmt.Println("Domain:  ", data["Domain"])
    fmt.Println("Method:  ", data["Method"])
    fmt.Println("Latency: ", data["Latency"])
}
```
###### OnSuccess is fired every time a service has check it be online

```go
func OnFailure(service map[string]interface{}) {
    fmt.Println("oh no! an event is failing right now! do something!")
    fmt.Println(service)
}
```
###### OnFailure is fired every time a service is failing

```go
func (p pkg) OnLoad(db sqlbuilder.Database) {
    fmt.Println("=============================================================")
    fmt.Printf("  Statup Example Plugin Loaded using %v database\n", db.Name())
    fmt.Println("=============================================================")
}
```
###### OnLoad is fired after plugin is loaded into the environment


## Interacting with Database
The Example Statup Plugin includes a variable `Database` that will allow you to interact with the Statup database. Checkout [database.go](https://github.com/hunterlong/statup_plugin/blob/master/database.go) to see a full example of Create, Read, Update and then Deleting a custom Communication entry into the database.
```go
// Insert a new communication into database
// once inserted, return the Communication
func (c *Communication) Create() *Communication {
	uuid, err := CommunicationTable().Insert(c)
	if err != nil {
		panic(err)
	}
	c.Id = uuid.(int64)
	return c
}
```

## Custom HTTP Routes
Plugin's can include their own HTTP route to accept GET/POST requests. Route are loaded after Statup loads all of it's Routes. Checkout [routes.go](https://github.com/hunterlong/statup_plugin/blob/master/routes.go) on the example plugin to see a full example of how to use it. 
```go
// You must have a Routes() method in your plugin
func (p *pkg) Routes() []plugin.Routing {
	return []plugin.Routing{{
		URL:     "hello",
		Method:  "GET",
		Handler: CustomInfoHandler,
	}}
}

// This is the HTTP handler for the '/hello' URL created above
func CustomInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Oh Wow!!! This is cool...")
}
```


## Plugin To-Do List
- [ ] Ability to includes assets like jpg, json, etc
