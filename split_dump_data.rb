puts "[+]start"

class String
  # camel to snake
  def underscore
    word = self.dup
    word.gsub!(/::/, '/')
    word.gsub!(/([A-Z]+)([A-Z][a-z])/,'\1_\2')
    word.gsub!(/([a-z\d])([A-Z])/,'\1_\2')
    word.tr!("-", "_")
    word.downcase!
    word
  end
end

# filepathをarrayに変換
def convert_filepath_to_array(path)
  lines = Hash.new([])
  c = 1
  File.open(path) do |file|
    file.each_line do |l|
      if l.include?("//+migu")
        c += 1
      end
      key = "key#{c}".to_sym
      lines[key] = lines[key] + [l]
    end
  end
  lines
end

# struct毎にfileに変換
def generate_gofile(lines, db_name)
  # import time必要かチェック
  disable_import_time = lines.select {|l| l.include?("time.Time")}.length.zero?
  # "type Page struct {" 文字列を "page"  に変換
  filename = lines.select {|l| l.include?("struct {")}[0].gsub("\n", '').gsub('type ', '').gsub(' struct {', '').underscore

  # 前処理
  path = "#{db_name}/#{filename}.go"
  body_header = "package #{db_name}\n"
  body_import = 'import "time"'

  # body生成
  if disable_import_time
    body = body_header + "\n"
    body += lines.join('')
  else
    body = body_header
    body += body_import + "\n\n"
    body += lines.join('')
  end

  # ファイルに書く
  Dir.mkdir(db_name, 0777) unless Dir.exist?(db_name)
  File.write(path, body)
  puts "generate: #{path}"
end

DB_NAME = 'matome'
PATH = 'dump.go'

# ファイル読む
lines = convert_filepath_to_array(PATH)

# gofileに出力
lines.keys.each do |key|
  next unless lines[key][0].include?("//+migu")
  generate_gofile(lines[key], DB_NAME)
end

puts "[+]finish"
