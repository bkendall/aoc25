load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "12_lib",
    srcs = ["main.go"],
    importpath = "aoc25/12",
    visibility = ["//visibility:private"],
    deps = ["//lib"],
)

go_binary(
    name = "12",
    embed = [":12_lib"],
    data = glob(["*.txt"]),
    visibility = ["//visibility:public"],
)

go_test(
    name = "12_test",
    size = "small",
    srcs = ["main_test.go"],
    embed = [":12_lib"],
    deps = ["@com_github_google_go_cmp//cmp"],
)
