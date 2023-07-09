<?php
function sanitizeInput($input) {
    // 过滤用户输入的内容，只允许部分 HTML 标签和属性
    $allowedTags = ['p', 'a', 'strong', 'em']; // 允许的标签
    $allowedAttributes = 'href,title'; // 允许的属性

    // 过滤 HTML 标签和属性
    $filteredContent = strip_tags($input, $allowedTags);
    $filteredContent = strip_tags($filteredContent, $allowedAttributes);

    // 检查过滤后的内容是否与原内容相同，如果不同则返回特定提示
    if ($input !== $filteredContent) {
        $disallowedTags = getDisallowedTags($input, $allowedTags);
        return $disallowedTags;
    }

    return $filteredContent;
}

function getDisallowedTags($input, $allowedTags) {
    $dom = new DOMDocument();
    libxml_use_internal_errors(true);
    $dom->loadHTML(mb_convert_encoding($input, 'HTML-ENTITIES', 'UTF-8'));
    $disallowedTags = [];

    foreach ($dom->getElementsByTagName('*') as $element) {
        $tag = $element->nodeName;
        if (!in_array($tag, $allowedTags)) {
            $disallowedTags[] = $tag;
        }
    }

    $disallowedTags = array_unique($disallowedTags);
    $disallowedTagsString = implode(', ', $disallowedTags);
    $disallowedTagsString = htmlspecialchars($disallowedTagsString);

    return $disallowedTagsString;
}

function sanitizeOutput($output) {
    // 转义输出的内容
    $escapedContent = htmlspecialchars($output);
    return $escapedContent;
}
?>
