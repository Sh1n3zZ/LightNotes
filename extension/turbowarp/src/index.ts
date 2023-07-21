import Extension from './include/plugin'
import axios from 'axios'

let show = false;

type Note = {
  id: number;
  title: string;
  body: string;
  updated_at: string;
}

type PaginationResponse = {
  status: boolean;
  total: number;
  page: number;
  prev_page: boolean;
  next_page: boolean;
  notes: Note[];
}

function setDisplay(state: boolean): boolean {
  show = state;
  const el = document.getElementById('notes');
  if (el) {
    el.className = show ? 'fade': '';
    el.style.zIndex = show ? '9999' : '-1';
  }
  return show;
}

function setToken(raw: string): void {
  const token = raw.split('').map((i: string) => {
    return String.fromCharCode(i.charCodeAt(0) + 1)
  }).join('');
  sessionStorage.setItem('notes', token);
  axios.defaults.headers.common['Authorization'] = raw;
}

function getToken(): string {
  const raw = sessionStorage.getItem('notes');
  if (!raw) return '';
  return raw.split('').map((i: string) => {
    return String.fromCharCode(i.charCodeAt(0) - 1)
  }).join('');
}

async function getNotes(page: number): Promise<{data: PaginationResponse, status: boolean, reason: string }> {
  try {
    const res = await axios.get(`/user/list?page=${page}`);
    if (!res.data.status) return { data: {} as PaginationResponse, status: false, reason: "接收失败！请检查您是否登录" };
    return { data: res.data as PaginationResponse, status: true, reason: "" };
  } catch (e) {
    console.debug(e);
    return { data: {} as PaginationResponse, status: false, reason: "接收失败！请检查您的网络环境" };
  }
}

new Extension({
  id: 'notes',
  name: '随记便签 Lightnotes',
  color1: '#69b7f7',
  blocks: [{
    opcode: "send",
    blockType: "reporter",
    text: "发送随记便签 标题 [title:string] 内容 [body:string] (漂流瓶特有的会丢失~)",
    default: { title: "新便签", body: "写点什么" },
    bind: async function ({ title, body }): Promise<string> {
      if (!title || !body) {
        return "标题或内容不能为空"
      }
      try {
        const res = await axios.post("/anonymous/send", { title, body });
        if (!res.data.status) return "发送失败！请稍后重试";
        return res.data.code;
      } catch (e) {
        console.debug(e);
        return "发送失败！请检查您的网络环境";
      }
    },
  }, {
    opcode: "get",
    blockType: "reporter",
    text: "获取随记便签 取签码 [code:number]",
    bind: async function ({ code }): Promise<string | Record<string, string>> {
      if (!code) {
        return "接收失败！取签码不能为空"
      }
      try {
        const res = await axios.get(`/anonymous/get?code=${code}`);
        if (!res.data.status) return "接收失败！请检查您的接签码是否正确，匿名便签是否过期";
        return JSON.stringify({
          title: res.data.title,
          body: res.data.body,
        })
      } catch (e) {
        console.debug(e);
        return "接收失败！请检查您的网络环境";
      }
    }
  }, {
    opcode: "exist",
    blockType: "Boolean",
    text: "随记便签是否存在 [code:number]",
    bind: async function ({ code }): Promise<boolean> {
      if (!code) return false
      try {
        const res = await axios.get(`/anonymous/get?code=${code}`);
        return res.data.status;
      } catch (e) {
        console.debug(e);
        return false;
      }
    }
  }, {
    opcode: "login",
    blockType: "command",
    text: "登录 lightnotes 账号",
    bind: function (): void {
      setDisplay(true);
    }
  }, {
    opcode: "close",
    blockType: "command",
    text: "关闭验证窗口",
    bind: function (): void {
      setDisplay(false);
    }
  }, {
    opcode: "logout",
    blockType: "command",
    text: "退出登录",
    bind: function (): void {
      sessionStorage.removeItem('notes');
      delete axios.defaults.headers.common['Authorization'];
    }
  }, {
    opcode: "check",
    blockType: "Boolean",
    text: "是否登录",
    bind: async function (): Promise<boolean> {
      if (!getToken()) return false;
      try {
        const res = await axios.post("/user/state");
        return res.data.status;
      } catch (e) {
        console.debug(e);
        return false;
      }
    }
  }, {
    opcode: "uget",
    blockType: "reporter",
    text: "获取便签  id [id:number]",
    bind: async function ({ id }): Promise<string | Record<string, string>> {
      try {
        const res = await axios.get(`/user/get?id=${id}`);
        if (!res.data.status) return "接收失败！请检查您的便签 id 是否正确";
        const note = res.data.note as Note;
        return JSON.stringify({
          title: note.title,
          body: note.body,
        });
      } catch (e) {
        console.debug(e);
        return "接收失败！请检查您的网络环境 以及是否登录！";
      }
    }
  }, {
    opcode: "utime",
    blockType: "reporter",
    text: "获取便签更新时间  id [id:number]",
    bind: async function ({ id }): Promise<string | number> {
      try {
        const res = await axios.get(`/user/get?id=${id}`);
        if (!res.data.status) return "接收失败！请检查您的便签 id 是否正确";
        return (res.data.note as Note).updated_at;
      } catch (e) {
        console.debug(e);
        return "接收失败！请检查您的网络环境 以及是否登录！";
      }
    }
  }, {
    opcode: "unew",
    blockType: "reporter",
    text: "新建便签 标题 [title:string] 内容 [body:string]",
    default: { title: "新便签", body: "写点什么" },
    bind: async function ({ title, body }): Promise<string | number> {
      if (!title || !body) {
        return "标题或内容不能为空"
      }
      try {
        const res = await axios.post("/user/save", { title, body });
        if (!res.data.status) return "发送失败！请稍后重试";
        return res.data.id;
      } catch (e) {
        console.debug(e);
        return "发送失败！请检查您的网络环境 以及是否登录！";
      }
    }
  }, {
    opcode: "uupdate",
    blockType: "reporter",
    text: "更新便签 id [id:number] 标题 [title:string] 内容 [body:string]",
    default: { title: "新便签", body: "写点什么" },
    bind: async function ({ id, title, body }): Promise<string | number> {
      if (!id || !title || !body) {
        return "标题或内容不能为空"
      }
      try {
        const res = await axios.post(`/user/update?id=${id}`, { id, title, body });
        if (!res.data.status) return "发送失败！请稍后重试";
        return res.data.status;
      } catch (e) {
        console.debug(e);
        return "发送失败！请检查您的网络环境 以及是否登录！";
      }
    }
  }, {
    opcode: "udelete",
    blockType: "reporter",
    text: "删除便签  id [id:number]",
    bind: async function ({ id }): Promise<string | number> {
      if (!id) {
        return "id 不能为空"
      }
      try {
        const res = await axios.post(`/user/delete?id=${id}`);
        if (!res.data.status) return "发送失败！请稍后重试";
        return res.data.status;
      } catch (e) {
        console.debug(e);
        return "发送失败！请检查您的网络环境 以及是否登录！";
      }
    }
  }, {
    opcode: "ulist",
    blockType: "reporter",
    text: "获取便签列表  第 [page:number] 页",
    default: { page: "1" },
    bind: async function ({ page }): Promise<string | Record<string, string>> {
      const resp = await getNotes(page);
      if (!resp.status) return resp.reason;
      return JSON.stringify(resp.data.notes);
    }
  }, {
    opcode: "utotal",
    blockType: "reporter",
    text: "获取便签页数",
    disableMonitor: true,
    bind: async function (): Promise<string | number> {
      const resp = await getNotes(1);
      if (!resp.status) return resp.reason;
      return resp.data.total;
    }
  }, {
    opcode: "uprev",
    blockType: "reporter",
    text: "第 [page:number] 页是否有上一页",
    default: { page: "1" },
    bind: async function ({ page }): Promise<string | boolean> {
      const resp = await getNotes(page);
      if (!resp.status) return resp.reason;
      return resp.data.prev_page;
    }
  }, {
    opcode: "unext",
    blockType: "reporter",
    text: "第 [page:number] 页是否有下一页",
    default: { page: "1" },
    bind: async function ({ page }): Promise<string | boolean> {
      const resp = await getNotes(page);
      if (!resp.status) return resp.reason;
      return resp.data.next_page;
    }
  }, {
    opcode: "website",
    blockType: "command",
    text: "Lightnotes 官网",
    bind: () => window.open("https://notes.lightxi.com"),
  }],
  i18n: {
    source: "zh",
    accept: ["zh", "en"],
  },
  onload: () => {
    const body = document.body;

    const animation = document.createElement('style');
    animation.innerHTML = `
    @keyframes fade-in {
      from {
        opacity: 0;
      }
      to {
        opacity: 1;
      }
    }
    .fade {
      animation: fade-in 0.3s ease-in-out forwards;
    }
    `;
    body.appendChild(animation);

    const container = document.createElement('div');
    container.id = 'notes';
    container.style.position = 'absolute';
    container.style.zIndex = '-1';
    container.style.overflow = 'hidden';
    container.style.top = '50%';
    container.style.left = '50%';
    container.style.transform = 'translate(-50%, -50%)';
    container.style.width = '60vh';
    container.style.height = '60vh';
    container.style.borderRadius = '12px';
    container.style.boxShadow = '0 0 12px rgba(0, 0, 0, 0.2)';
    container.style.opacity = '0';
    body.appendChild(container);

    const close = document.createElement('svg');
    close.style.position = 'absolute';
    close.style.top = '0';
    close.style.right = '0';
    close.style.width = '24px';
    close.style.height = '24px';
    close.style.margin = '6px';
    close.style.cursor = 'pointer';
    close.innerHTML = '<path d="m16.192 6.344-4.243 4.242-4.242-4.242-1.414 1.414L10.535 12l-4.242 4.242 1.414 1.414 4.242-4.242 4.243 4.242 1.414-1.414L13.364 12l4.242-4.242z">'
    close.onclick = () => {
      body.removeChild(container);
    }
    container.appendChild(close);

    const frame = document.createElement('iframe');
    frame.src = 'https://notes.lightxi.com/';
    frame.style.width = '100%';
    frame.style.height = '100%';
    frame.style.border = 'none';
    frame.onload = () => {
      const iframe = frame.contentWindow;
      if (!iframe) return;
      setInterval(() => {
        iframe.postMessage({ type: 'login' }, '*');
      }, 1000);
      window.addEventListener('message', (e) => {
        if (e.data.type === 'login') {
          if (e.data.status) {
            setToken(e.data.token);
            setDisplay(false);
          }
        }
      }, false);
    }
    container.appendChild(frame);

    axios.defaults.baseURL = 'https://notes.lightxi.com/api';
    axios.defaults.headers.common['Authorization'] = getToken();
    axios.defaults.headers.common['Content-Type'] = 'application/json';
  }
}).register()
