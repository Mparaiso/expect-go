expect-go
---------

A helper function for go tests.

copyrights: 2016 mparaiso <mparaiso@online.fr>
license: GPL-3.0


#### Installation :

    go get github.com/mparaiso/expect-go 

#### Usage:

    package main_testgit 

    import "testing"
    import "github.com/mparaiso/expect-go" 

    func Test(t *testing.T){
        // expects "want" to equal "got" (which will yield a test error)
        expect.Expect(t,"want","got","example")
    }