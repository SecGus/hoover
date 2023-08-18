# hoover (note this readme is out of date and needs to be addressed)
Simple tool to remove all files that match an md5sum.

```
Usage of ./hoover:
  -d string
        directory of files to check (default: cwd) (default "./")
  -f string
        Required: example bad image file (example: 404_image.jpeg)
  -h string
        md5sum of example bad image file
  -s string
        Flag to silence the binary
```

Example:
```
chivato@kingdom:~$ cp hoover/deleteme ./file_to_be_deleted
chivato@kingdom:~$ cd hoover/
chivato@kingdom:~/hoover$ ./hoover -f deleteme -d '../'

                  ||
                  ||
                  ||
                  ||
                  ||
                  ||
                  ||     Here you go, sweep
                  ||     that up..............
                 /||\
                /||||\
                ======         __|__
                ||||||        / ~@~ \
                ||||||       |-------|
                ||||||       |_______|

Deleting file ../file_to_be_deleted
chivato@kingdom:~/hoover$
```
