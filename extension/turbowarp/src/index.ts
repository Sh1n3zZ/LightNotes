import Extension from './include/plugin'
import axios from 'axios'

let show = false;
let offsetX: number, offsetY: number, isDragging = false;

function setDisplay(state: boolean): boolean {
  show = state;
  const el = document.getElementById('notes');
  if (el) el.className = show ? 'fade': '';
  return show;
}

new Extension({
  id: 'notes',
  name: '随记便签',
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
        const res = await axios.post("https://notes.lightxi.com/api/anonymous/send", { title, body }, {
          headers: {
            'Content-Type': 'application/json'
          },
        })
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
        const res = await axios.get(`https://notes.lightxi.com/api/anonymous/get?code=${code}`, {
          headers: {
            'Content-Type': 'application/json'
          },
        });
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
    text: "便签是否存在 [code:number]",
    bind: async function ({ code }): Promise<boolean> {
      if (!code) {
        return false
      }
      try {
        const res = await axios.get(`https://notes.lightxi.com/api/anonymous/get?code=${code}`, {
          headers: {
            'Content-Type': 'application/json'
          }
        });
        return res.data.status;
      } catch (e) {
        console.debug(e);
        return false;
      }
    }
  }, {
    opcode: "login",
    blockType: "command",
    text: "登录随记便签 账号",
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
    opcode: "website",
    blockType: "command",
    text: "寻找河流的源头...",
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
    container.style.zIndex = '1024';
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
      iframe.postMessage({ type: 'login' }, '*');
    }
    container.appendChild(frame);

    container.addEventListener('mousedown', (event) => {
      isDragging = true;
      offsetX = event.clientX - container.offsetLeft;
      offsetY = event.clientY - container.offsetTop;
      container.style.cursor = 'grabbing';
    });

    document.addEventListener('mousemove', (event) => {
      if (!isDragging) return;
      const x = event.clientX - offsetX;
      const y = event.clientY - offsetY;
      container.style.left = x + 'px';
      container.style.top = y + 'px';
    });

    document.addEventListener('mouseup', () => {
      isDragging = false;
      container.style.cursor = 'grab';
    });
  }
}).register()
