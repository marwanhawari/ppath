class Ppath < Formula
  desc "🌈 A command-line tool to pretty print your system's PATH environment variable."
  homepage "https://github.com/marwanhawari/ppath"
  head "https://github.com/marwanhawari/ppath.git", branch: "main"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w")
  end

  test do
    system bin/"ppath", "-h"
  end
end
