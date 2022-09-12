Simple random files generation tool and read speed benchmark

Uses https://github.com/hellmany/hbench package

    ./hbench-cli_linux_amd64 -h
    Usage of ./hbench-cli_linux_amd64:
      -a string
            Action gen|bench - required
      -i int
            interations to read (default 1)
      -m int
            max files to read/wrie, 0 - to read all (default 1000)
      -p string
            path (default "/stor/tmp")
      -r int
            add rand mb to size writing (default 3)
      -s int
            size to read/write in Mb, 0 - full file to read (default 1)
      -t int
            threads to read/write (default 100)
Sample output

        2022/09/11 07:36:12 Writed 20000 files and 79959.00 megabytes  ( 78.08 gigabytes ) in 100 threads
    2022/09/11 07:36:12 Speed: 247.79 mb/s
    2022/09/11 07:36:12 Took 5m22.684907641s ( 322.684907641 ) seconds
        
    
    2022/09/11 07:41:04 Readed 10000 files and 50849.00 megabytes  ( 49.66 gigabytes ) in 100 threads
    2022/09/11 07:41:04 Speed: 248.89 mb/s
    2022/09/11 07:41:04 Took 3m24.30025652s ( 204.30025652 ) seconds
    
    
