# d2-item-sorter

Sort your diablo2 Plugy Stash automatically based on a configuration file.

## Disclaimer

This is **WORK IN PROGRESS**! Please consider this **EXPERIMENTAL** in the current stage and use it only at your own risk. I do not take any responsibility for any damage caused by this.

## Usage

At the moment only tested with D2 LOD 1.13 and Plugy 14. 

**Please backup your Save folder before using this! Altough I'm using this on my own I do not take any responsibility for any lost or damaged save files!** 

```
>d2-item-sorter.exe help

Start Server to access UI on http://localhost:port

Options:

  -h, --help          display help information
      --port[=3000]   specify port number for server to listen on

Commands:

  help       display help information
  watch      Watches the shared stash to record runs and automatically sort items after each run
  transfer   Transfers items from one character to another or to the shared stash (use _shared instead of a character name)
  info       Lists Information about the items you have
  sort       Sorts Items of a character or the shared stash
```

### Sorting

The commands (transfer, sort) will automatically sort your items based on the configuration found in `./groupConfigs/groups.yaml` and create pages based on this configuration. This means that all items will be ordered based on the configuration. 

_Please be aware that in the current version the transfer command does not work in every scenario and duplicated stash pages might be the result - this is a known issue at the moment._

### Automatic Sorting

 When starting d2-item-sorter with the `watch`command it will watch the shared stash save file for changes and after every run you do it will create a backup sort your items, record a run, track the time between the last run and also print out which items you have newly added (compared to the last run). 
This can be really helpful in comparison with the pickup page mentioned below.

#### Ignore Page during sorting

If you want to keep a page unsorted (maybe to keep special items prepared for another character or friend) you can just rename the page and include `(i)` in front of the name (i.e `"(i) Items for Friend"`). The sorting algorithm will then **ignore** this page and the items on it.


#### Pickup Page

There is also the posibility to prefix a page name with `(p)` which stands for **pickup**. The items on this page will get sorted but the page will not be overwritten during the sorting. A use case would be to declare the first (or any) page in your shared stash as pickup page. Move items there after every run and they will be picked up, get sorted but the page will remain


## TODOS:
- [x] get size of all items
- [x] store item position on page
- [ ] arrange items in pages
- [x] check collisions
- [x] change position in item
- [x] write pages to file
- [x] stash page flags (flag for shared stash, and also for different index types)
- [x] bit writer
- [ ] web ui to show shared and character stashes
- [ ] holy grail status

## Known Issues:
When transfering items from a character to the shared stash there is a bug which does not write the character file correctly and so the currently open page will be duplicated. 

Please report any issues you have in github I will take a look at it whenever time admits.

## Compile and Start

Make sure that the folder groupConfigs is in the same folder as your .exe file
```
go build -o ./dist/d2-item-sorter.exe . && .\dist\d2-item-sorter.exe
```

# Credits

The following projects helped me getting this started and also provided very valuable insights in the d2 file structure. Thank you guys for what you have provided with your great projects.

- [nokka/d2s](https://github.com/nokka/d2s) I'm using a forked version of this which exports some additional information to be used by d2-item-sorter. After I have a stable base I will propose some changes to d2s and maybe the fork will then not be needed anymore
- [mrpara/d2_plugy_stash_organizer](https://github.com/mrpara/d2_plugy_stash_organizer). This really gave me the idea to start, but as I found d2s and wanted to learn go I took some ideas (especially the collision part) and ported it to this go project
- https://squeek502.github.io/d2itemreader/ Very good insights in the d2 file formats, helped me a lot!
