# CLI Bug tracker
Bug Tracker made in Go using Cobra.

[![Is compiling?](https://github.com/faculerena/bugtracker/actions/workflows/main.yml/badge.svg)](https://github.com/faculerena/bugtracker/actions/workflows/main.yml)
# About the tracker
Small CLI app to track bugs, you can:

* **ADD** new bugs to track
* **GET \<ID>** returns that bug.
* Mark as **SOLVED \<ID>** an existing bug
* Create a **NOTE** for an existing bug **[WIP]**
* **DELETE \<ID>** to delete a bug.
* **CLEAR** the tracker
* **RELATE \<ID> <ID_target>** a bug with another.
  * You can use **RELATED \<ID>** to retrieve all the bugs related to an ID (including itself)
* **REOPEN** a solved bug
* **LIST** to retrieve all open bugs. 
* **EDIT \<ID>** an existing bug 

# Save
For now, all the bugs are stored in .json format in home dir in a file named ".tracker.json"
The next usable id is saved there too on a ".id" file. I will make this a better way.


# How to install

With Go ver 1.19.5 you can:

Build the file

``go build``


Then run it as 

```./tracker [cmd]```


# Images

```add``` asks for "what", "how", and "priority"

![add](readmeImages/add.png)

```list``` retrieves unsolved bugs

![list](readmeImages/list.png)

```list all``` retrieves all bugs

![list all](readmeImages/listall.png) 

