
Currently Tue 24 Jul 2018 21:52:55 AEST 


sudo apt install golang-go
prepares to install : golang-1.6-go golang-1.6-src golang-src
Note : version 1.6 so don't do this !


The download page under golang.org helps you to download version 1.10.3
 (Typically these commands must be run as root or through sudo.)

Add /usr/local/go/bin to the PATH environment variable. You can do this by adding this line to your /etc/profile (for a system-wide installation) or $HOME/.profile:

export PATH=$PATH:/usr/local/go/bin


The instructions at https://golang.org/doc/install?download=go1.10.3.linux-amd64.tar.gz    are, for me, a bit confusing, so I have created some instructions for myself...






