import Extension from './include/plugin'

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
        const res = await fetch("https://notes.lightxi.com/api/anonymous/send", {
          method: "POST",
          headers: {
            "Content-Type": "application/json"
          },
          body: JSON.stringify({ title, body }),
        });
        const data = await res.json();
        if (!data.status) return "发送失败！请稍后重试";
        return data.code;
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
        const res = await fetch(`https://notes.lightxi.com/api/anonymous/get?code=${code}`);
        const data = await res.json();
        if (!data.status) return "接收失败！请检查您的接签码是否正确，匿名便签是否过期";
        return {
          title: data.title,
          body: data.body,
        };
      } catch (e) {
        console.debug(e);
        return "接收失败！请检查您的网络环境";
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
