#!/usr/bin/env ruby

system 'git branch -a'

puts `git branch -a`

require 'open3'

Open3.popen2('git branch -a') do |input, output, wait_t|
  until output.eof?
    puts output.gets
  end
end
