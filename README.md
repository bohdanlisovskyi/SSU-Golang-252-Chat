# SSU-Golang-252-Chat

Repository contains sources for server and client application.

Each application stores in separate folder and has its own Makefile.

### How to start
Run to install glide and gometalinter

    make install-helpers

Run toinstall application dependencies

    make dependencies
    
### Code quality check
    make code-quality

### How to install GUI library and run build
Follow the instruction https://github.com/therecipe/qt#installation . If needed see "Regular Setup"

To build application run
 
    qtdeploy build desktop [path/to/your/project] 

or 
    
    $GOPATH/bin/qtdeploy build desktop [path/to/your/project]