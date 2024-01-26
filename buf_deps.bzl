load("@rules_buf//buf:defs.bzl", "buf_dependencies")

def buf_deps():
    buf_dependencies(
        name = "buf_deps",
        modules = [
            "buf.build/googleapis/googleapis:a86849a25cc04f4dbe9b15ddddfbc488",
        ],
    )
