dosa-idl
========
[DOSA](https://github.com/uber-go/dosa/) - Declarative Object Storage Abstraction

This repo holds thrift IDL files and generated code used by DOSA to communicate between server and clients. This repo is not end-user useful on its own, but needed only when a feature require IDL changes.

thriftrw
--------

Install thriftrw:

    go get -u go.uber.org/thriftrw
    cd .../thriftrw
    git checkout v1.19.0

Use the appropriate version.

License
-------
MIT License, please see [LICENSE](https://github.com/uber/dosa-idl/blob/master/LICENSE) for details.

