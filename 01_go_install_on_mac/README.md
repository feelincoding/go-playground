# How to install go on Mac(m1)

1. brew update
2. brew install go
3. create go file, `hello.go`
4. if you see message that 'the "go-outline" command is not available ~', should click to `install-all` button.
5. You can see out put
    ```
    Tools environment: GOPATH=/Users/ijeongpil/go
    Installing 7 tools at /Users/ijeongpil/go/bin in module mode.
    gotests
    gomodifytags
    impl
    goplay
    staticcheck
    gopls
    go-outline

    Installing github.com/cweill/gotests/gotests@latest (/Users/ijeongpil/go/bin/gotests) SUCCEEDED
    Installing github.com/fatih/gomodifytags@latest (/Users/ijeongpil/go/bin/gomodifytags) SUCCEEDED
    Installing github.com/josharian/impl@latest (/Users/ijeongpil/go/bin/impl) SUCCEEDED
    Installing github.com/haya14busa/goplay/cmd/goplay@latest (/Users/ijeongpil/go/bin/goplay) SUCCEEDED
    Installing honnef.co/go/tools/cmd/staticcheck@latest (/Users/ijeongpil/go/bin/staticcheck) SUCCEEDED
    Installing golang.org/x/tools/gopls@latest (/Users/ijeongpil/go/bin/gopls) SUCCEEDED
    Installing github.com/ramya-rao-a/go-outline@latest (/Users/ijeongpil/go/bin/go-outline) SUCCEEDED

    All tools successfully installed. You are ready to Go. :)
    ```
6. Code it like `hello.go` file