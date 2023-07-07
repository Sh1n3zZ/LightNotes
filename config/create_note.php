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
                                noteList.append("<p>" + note.content + "</p>");
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
</body>
</html>
