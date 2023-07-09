<!DOCTYPE html>
<html>
<head>
    <title>创建便签</title>
    <style>
        .pagination-link {
            display: inline-block;
            margin-right: 5px;
            padding: 3px 6px;
            cursor: pointer;
            border: 1px solid #ccc;
            border-radius: 3px;
        }
        .active {
            background-color: #ccc;
        }
    </style>
        <style>
        /* Markdown 解析后的样式 */
        .note p,
        .note h1,
        .note h2,
        .note h3,
        .note h4,
        .note h5,
        .note h6 {
            margin: 0 0 10px;
        }

        .note p {
            line-height: 1.5;
        }

        .note h1,
        .note h2,
        .note h3,
        .note h4,
        .note h5,
        .note h6 {
            font-weight: bold;
        }

        .note h1 {
            font-size: 24px;
        }

        .note h2 {
            font-size: 20px;
        }

        .note h3 {
            font-size: 18px;
        }

        .note h4 {
            font-size: 16px;
        }

        .note h5 {
            font-size: 14px;
        }

        .note h6 {
            font-size: 12px;
        }
    </style>
    <script src="https://cdn.jsdelivr.net/npm/marked@2.1.3/marked.min.js"></script>
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.6.0/jquery.min.js"></script>
    <script>
        $(document).ready(function() {
            var currentPage = 1;
            var notesPerPage = 5;

            loadNotes(currentPage);

            $(document).on("click", ".pagination-link", function() {
                currentPage = $(this).data("page");
                loadNotes(currentPage);
            });

            function loadNotes(page) {
                $.ajax({
                    url: "get_notes.php",
                    type: "GET",
                    dataType: "json",
                    data: { page: page, perPage: notesPerPage },
                    success: function(data) {
                        if (data.notes.length > 0) {
                            var noteList = $("#note-list");
                            noteList.empty();

                            $.each(data.notes, function(index, note) {
                                var content = unescape(note.content); // 解码含有 Unicode 转义字符的内容
                                var html = marked(content); // 解析 Markdown 文本为 HTML
                                noteList.append("<div class='note'>" + html + "</div>");
                            });

                            var pagination = $("#pagination");
                            pagination.empty();

                            if (data.totalPages > 1) {
                                for (var i = 1; i <= data.totalPages; i++) {
                                    if (i === page) {
                                        pagination.append("<span class='pagination-link active' data-page='" + i + "'>" + i + "</span>");
                                    } else {
                                        pagination.append("<span class='pagination-link' data-page='" + i + "'>" + i + "</span>");
                                    }
                                }
                            }
                        } else {
                            $("#note-list").text("没有便签可展示");
                        }
                    },
                    error: function() {
                        $("#note-list").text("加载便签失败");
                    }
                });
            }
        });
    </script>
</head>
<body>
    <h2>创建便签</h2>
    <form action="save_note.php" method="POST">
        <label for="note_content">便签内容：</label>
        <textarea id="note_content" name="note_content" rows="5" cols="30" required></textarea><br><br>
        <input type="submit" value="保存便签">
    </form>

    <h2>我的便签</h2>
    <div id="note-list"></div>
    <div id="pagination"></div>
</body>
</html>
