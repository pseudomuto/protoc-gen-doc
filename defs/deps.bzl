load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "com_github_aokoli_goutils",
        importpath = "github.com/aokoli/goutils",
        sum = "h1:7fpzNGoJ3VA8qcrm++XEE1QUe0mIwNeLa02Nwq7RDkg=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_davecgh_go_spew",
        importpath = "github.com/davecgh/go-spew",
        sum = "h1:XxMZvQZtTXpWMNWK82vdjCLCe7uGMFXdTsJH0v3Hkvw=",
        version = "v0.0.0-20161028175848-04cdfd42973b",
    )
    go_repository(
        name = "com_github_envoyproxy_protoc_gen_validate",
        importpath = "github.com/envoyproxy/protoc-gen-validate",
        sum = "h1:bV5JGEB1ouEzZa0hgVDFFiClrUEuGWRaAc/3mxR2QK0=",
        version = "v0.3.0-java",
    )
    go_repository(
        name = "com_github_gogo_protobuf",
        importpath = "github.com/gogo/protobuf",
        sum = "h1:72R+M5VuhED/KujmZVcIquuo8mBgX4oVda//DQb3PXo=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_golang_protobuf",
        importpath = "github.com/golang/protobuf",
        sum = "h1:0iH4Ffd/meGoXqF2lSAhZHt8X+cPgkfn/cb6Cce5Vpc=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_google_uuid",
        importpath = "github.com/google/uuid",
        sum = "h1:jWtZjFEUE/Bz0IeIhqCnyZ3HG6KRXSntXe4SjtuTH7c=",
        version = "v0.0.0-20161128191214-064e2069ce9c",
    )
    go_repository(
        name = "com_github_huandu_xstrings",
        importpath = "github.com/huandu/xstrings",
        sum = "h1:pO2K/gKgKaat5LdpAhxhluX2GPQMaI3W5FUz/I/UnWk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_imdario_mergo",
        importpath = "github.com/imdario/mergo",
        sum = "h1:mKkfHkZWD8dC7WxKx3N9WCF0Y+dLau45704YQmY6H94=",
        version = "v0.3.4",
    )
    go_repository(
        name = "com_github_masterminds_semver",
        importpath = "github.com/Masterminds/semver",
        sum = "h1:WBLTQ37jOCzSLtXNdoo8bNM8876KhNqOKvrlGITgsTc=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_masterminds_sprig",
        importpath = "github.com/Masterminds/sprig",
        sum = "h1:0gSxPGWS9PAr7U2NsQ2YQg6juRDINkUyuvbb4b2Xm8w=",
        version = "v2.15.0+incompatible",
    )
    go_repository(
        name = "com_github_mwitkow_go_proto_validators",
        importpath = "github.com/mwitkow/go-proto-validators",
        sum = "h1:28i1IjGcx8AofiB4N3q5Yls55VEaitzuEPkFJEVgGkA=",
        version = "v0.0.0-20180403085117-0950a7990007",
    )
    go_repository(
        name = "com_github_pmezard_go_difflib",
        importpath = "github.com/pmezard/go-difflib",
        sum = "h1:GD+A8+e+wFkqje55/2fOVnZPkoDIu1VooBWfNrnY8Uo=",
        version = "v0.0.0-20151028094244-d8ed2627bdf0",
    )
    go_repository(
        name = "com_github_pseudomuto_protokit",
        importpath = "github.com/pseudomuto/protokit",
        sum = "h1:hlnBDcy3YEDXH7kc9gV+NLaN0cDzhDvD1s7Y6FZ8RpM=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        sum = "h1:Zx8Rp9ozC4FPFxfEKRSUu8+Ay3sZxEUZ7JrCWMbGgvE=",
        version = "v0.0.0-20170130113145-4d4bfba8f1d1",
    )
    go_repository(
        name = "in_gopkg_check_v1",
        importpath = "gopkg.in/check.v1",
        sum = "h1:yhCVgyC4o1eVCa2tZl7eS0r+SDo693bJlVdllGtEeKM=",
        version = "v0.0.0-20161208181325-20d25e280405",
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        importpath = "gopkg.in/yaml.v2",
        sum = "h1:ZCJp+EgiOT7lHqUV2J862kp8Qj64Jo6az82+3Td9dZw=",
        version = "v2.2.2",
    )
    go_repository(
        name = "org_golang_google_genproto",
        importpath = "google.golang.org/genproto",
        sum = "h1:b69RmkJsx8NyRJsKF2mQ/AF8s4BNxwNsT4rQ3wON1U0=",
        version = "v0.0.0-20181107211654-5fc9ac540362",
    )
    go_repository(
        name = "org_golang_x_crypto",
        importpath = "golang.org/x/crypto",
        sum = "h1:O5C+XK++apFo5B+Vq4ujc/LkLwHxg9fDdgjgoIikBdA=",
        version = "v0.0.0-20180501155221-613d6eafa307",
    )
    go_repository(
        name = "org_golang_x_sync",
        importpath = "golang.org/x/sync",
        sum = "h1:IqXQ59gzdXv58Jmm2xn0tSOR9i6HqroaOFRQ3wR/dJQ=",
        version = "v0.0.0-20190412183630-56d357773e84",
    )
