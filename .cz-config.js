'use strict';

module.exports = {

  types: [
    {
      value: 'WIP',
      name : 'WIP: 暫存工作。'
    },
    {
      value: 'feat',
      name : 'feat: 新增功能。'
    },
    {
      value: 'docs',
      name : 'docs: 修改文件。'
    },
    {
      value: 'style',
      name : 'style: 修改程式碼的風格，不會對產品有任何的功能變動 (空白鍵、格式化、分號...等)。'
    },
    {
      value: 'revert',
      name : 'revert: 撤銷、復原一次 git commit。'
    }
  ],
  messages: {
    type: '<type> 用於說明 commit 的類別，只允許使用下面 5 個標識: \n',
    scope: '<scope> 影響範圍，比如會影響到哪個模塊/性能/哪一層（業務層，持久層，緩存，rpc），如果是特性代碼，可以寫特性名稱 (可選): \n',
    customScope: '<scope> 自定義影響範圍，請精簡扼要但不失原意: \n',
    subject: '<subject> 目的的簡短描述，不超過 100 個字符: \n',
    body: '<body> 對本次 commit 的詳細描述，使用第一人稱，應該說明代碼變動的動機，以及與以前行為的對比，可以使用 "|" 分成多行 (可選):\n',
    breaking: '<breaking> 對破壞性變動(Breaking Change)的描述、以及變動理由和遷移方法 (可選):\n',
    footer: '<footer> 針對的 issue，像是：#520, #1314 (可選):\n',
    confirmCommit: '<confirm commit> 請確認以上描述。',
  },
  upperCaseSubject: true,
  footerPrefix: 'Related issue:',
  allowCustomScopes: true,
  allowBreakingChanges: ["feat", "fix", "refactor"],
};