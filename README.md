# Statup Plugin Example
An example of a [Statup](https://github.com/hunterlong/statup) plugin using the Golang plugins feature. Remember to build this plugin on the same operating system type as 
your Statup service is running. To build your own plugins, checkout the [Statup Plugins Wiki](https://github.com/hunterlong/statup/wiki/Statup-Plugins) page for detailed information about this process.

## Building
Clone this repo, and run the commands below...
```bash
git clone https://github.com/hunterlong/statup_plugin
cd statup-plugin

go build -buildmode=plugin -o example.so
```
Once you have `example.so`, drag it into the `plugins` folder in your Statup directory. Reload Statup service and it should 
have some text for `OnLoad()`! 

<p align="center">
<img width="95%" src="https://s3-us-west-2.amazonaws.com/gitimgs/statupplugin.png">
</p>

## Testing
This is an example for a Statup plugin that interacts with HTTP Routes, and the Statup Database. To test this plugin, or your plugin, 
insert the compiled `.so` file into your `plugins` directory and then run:
```bash
statup test plugins
```
This command will create a temporary database and triggers events to your plugin.

## Creating a Custom Plugin
The plan is to give users the ability to create their own plugins, public and private. This repo will be updated when there 
are more features and changes to the Statup plugin package. Enjoy!
