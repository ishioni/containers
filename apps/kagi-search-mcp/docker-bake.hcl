target "docker-metadata-action" {}

variable "APP" {
  default = "kagi-search-mcp"
}

variable "VERSION" {
  // renovate: datasource=pypi depName=kagimcp versioning=pep440
  default = "0.1.5"
}

variable "SOURCE" {
  default = "https://github.com/kagisearch/kagimcp"
}

group "default" {
  targets = ["image-local"]
}

target "image" {
  inherits = ["docker-metadata-action"]
  args = {
    VERSION = "${VERSION}"
  }
  labels = {
    "org.opencontainers.image.source" = "${SOURCE}"
  }
}

target "image-local" {
  inherits = ["image"]
  output = ["type=docker"]
  tags = ["${APP}:${VERSION}"]
}

target "image-all" {
  inherits = ["image"]
  platforms = [
    "linux/amd64",
    "linux/arm64"
  ]
}
