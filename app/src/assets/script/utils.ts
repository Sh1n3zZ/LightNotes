export function formatDate(time: string | Date, offset: boolean = true): string {
  const now = new Date(), date = typeof time == 'string' ? new Date(time) : time;
  const diff = (now.getTime() - date.getTime()) / 1000 + (offset ? 8 * 3600 : 0); // second

  if (diff < 0) {
    return '无';
  } else if (diff < 60) {
    return '刚刚';
  } else if (diff < 3600) {
    const minutes = Math.floor(diff / 60);
    return `${minutes} 分钟前`;
  } else if (diff < 86400) {
    const hours = Math.floor(diff / 3600);
    return `${hours} 小时前`;
  } else if (diff < 172800) {
    return `昨天 ${padZero(date.getHours())}:${padZero(date.getMinutes())}`;
  } else if (diff < 259200) {
    return `前天 ${padZero(date.getHours())}:${padZero(date.getMinutes())}`;
  } else if (diff < 604800) {
    const days = Math.floor(diff / 86400);
    return `${days} 天前 ${padZero(date.getHours())}:${padZero(date.getMinutes())}`;
  } else if (date.getFullYear() === now.getFullYear()) {
    return `${date.getMonth() + 1}/${date.getDate()} ${padZero(date.getHours())}:${padZero(date.getMinutes())}`;
  } else {
    return `${date.getFullYear()}/${date.getMonth() + 1}/${date.getDate()} ${padZero(date.getHours())}:${padZero(date.getMinutes())}`;
  }
}

export function padZero(n: number): string {
  return (n < 10 ? '0' : '') + n;
}

export function contain(el: HTMLElement | null | undefined, target: HTMLElement | null): boolean {
  return (el && target) ? (el == target || el.contains(target)) : false;
}
