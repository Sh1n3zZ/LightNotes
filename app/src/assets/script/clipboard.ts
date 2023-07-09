export async function copyText(text: string): Promise<boolean> {
  if (!navigator.clipboard) {
    const input = document.createElement("input");
    input.style.display = "none";
    input.value = text;
    document.body.appendChild(input);
    input.select();
    document.execCommand("copy");
    document.body.removeChild(input);
    return true;
  }

  try {
    await navigator.clipboard.writeText(text);
    return true;
  } catch (err) {
    return false;
  }
}
