# Description
I hate all the manga/comic reading apps for Linux, they all suck and the only good one is mangareader. The one I have in windows is very nice and organizes everything pretty well.
this will be the start of a terminal based manga reader using golang.

This tui application main focus will be to just be used only for local collection.

# TODO:
- [ ] implement search/filter component
- [ ] implement file tree/picker component
- [ ] get familiar with the bubble tea framework: understand how to handle state, and also render app dimensions
- [ ] understand how we can render images on the terminal
- [ ] research alternative paths just incase tui idea is not possible with golang

# App desired functionality
App should allow for more than one filepath for local collection
- [ ] Search and filter local collection.
- [ ] have collection separated by which filepath is imported
## Example
```
Main/file/path/tree1
├── series-subtree-1/
│   ├── somecbz1
│   ├── somecbz2
│   ├── somecbz3
├── series/subtree2/
│   ├── somecbz1
│   ├── somecbz2
│   ├── somecbz3

Main/file/path/tree2
├── series/subtree1/
│   ├── somecbz1
│   ├── somecbz2
│   ├── somecbz3
├── series/subtree2/
│   ├── somecbz1
│   ├── somecbz2
│   ├── somecbz3
```
- [ ] image rendering will be the hardest so look into that
- [ ] if possible allow an image preview while selecting something to read

# Prototype Visualization
![prototype](./docs/mangatuidesignidea.png)
