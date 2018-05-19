# dump
migu dump -h 127.0.0.1 -u root matome ./d

# dumpしたファイルを分離
ruby ./split_dump_data.rb

# dry run
migu sync -h 127.0.0.1 -u root --dry-run matome dump.go
migu sync -h 127.0.0.1 -u root --dry-run matome ./matome

# sync
migu sync -h 127.0.0.1 -u root matome dump.go
