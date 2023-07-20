import Extension from './include/plugin'
import axios from 'axios'

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
    opcode: "website",
    blockType: "command",
    text: "寻找河流的源头...",
    bind: () => window.open("https://notes.lightxi.com"),
  }],
  i18n: {
    source: "zh",
    accept: ["zh", "en"],
  }
}).register()
