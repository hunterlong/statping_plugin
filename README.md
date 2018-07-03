# Statup Plugin Example
An example of a [Statup](https://github.com/hunterlong/statup) plugin using the Golang plugins feature. Remember to build this plugin on the same operating system type as 
your Statup service is running.

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

## Creating a Custom Plugin
The plan is to give users the ability to create their own plugins, public and private. This repo will be updated when there 
are more features and changes to the Statup plugin package. Enjoy!
