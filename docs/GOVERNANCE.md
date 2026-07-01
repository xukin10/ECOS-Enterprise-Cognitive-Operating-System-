# AetherOS Governance Framework

本文档描述 AetherOS 项目如何使用 GitHub Projects、Discussions、RFC 和 ADR 进行决策和规划。

## 🎯 目标

- 透明、可追踪的决策过程
- 社区参与和反馈
- 完整的决策历史记录
- 清晰的实施路线图

---

## 📋 工作流程

### 1️⃣ 发现阶段 - GitHub Discussions

**何时使用：** 初期想法、社区反馈、头脑风暴

**操作步骤：**
- 在 **Discussions** 中发起新话题
- 标签分类：`idea`、`design`、`question`、`feedback`
- 收集社区意见和建议
- 达成共识后，推进到 RFC 阶段

**示例讨论标题：**
- "Should we adopt X architecture?"
- "How should we handle Y scenario?"

---

### 2️⃣ 提案阶段 - RFC (Request for Comments)

**何时使用：** 重大功能、架构变更、重要决策

**操作步骤：**
1. 创建 RFC 文件：`docs/rfcs/RFC-XXX-[title].md`
2. 使用 RFC 模板（见 `RFC-TEMPLATE.md`）
3. 创建 Pull Request 到 `main` 分支
4. PR 标记为 `type:rfc`
5. 收集评论和建议
6. 迭代更新设计
7. 最终标记为 Accepted/Rejected

**RFC 命名规则：**
```
RFC-001-async-event-system.md
RFC-002-plugin-architecture.md
```

**PR 描述模板：**
```markdown
## RFC: [Title]

相关讨论（如有）：#XXX

### 目的
简述为什么需要这个提案。

### 关键决策点
- [ ] 架构设计
- [ ] 性能影响
- [ ] 向后兼容性

### 等待反馈
请在下方评论您的看法。
```

---

### 3️⃣ 决策阶段 - ADR (Architecture Decision Record)

**何时使用：** RFC 被接受后，记录最终决策

**操作步骤：**
1. 创建 ADR 文件：`docs/adrs/ADR-XXX-[title].md`
2. 使用 ADR 模板（见 `ADR-TEMPLATE.md`）
3. 引用相关的 RFC 和 Discussion
4. 标记状态为 `Accepted`
5. 合并到 `main` 分支

**ADR 命名规则：**
```
ADR-001-async-event-system.md
ADR-002-plugin-architecture.md
```

**ADR 生命周期：**
- `Proposed` - 初次提案
- `Accepted` - 已决策，准备实施
- `Deprecated` - 仍有效但不推荐
- `Superseded` - 被更新的 ADR 替代

---

### 4️⃣ 执行阶段 - GitHub Projects

**何时使用：** 追踪实施工作、管理里程碑

**操作步骤：**
1. 在 Projects 中创建相应卡片
2. 卡片链接到对应的 ADR/RFC
3. 拆分为具体 Issue
4. 追踪进度：Backlog → In Progress → In Review → Done

**Project 看板列：**
- Backlog - 待规划工作
- In Progress - 正在进行
- In Review - 代码审查中
- Done - 已完成

---

## 🔗 文档关系图

```
Discussion (社区反馈)
    ↓ (达成共识)
RFC Pull Request (设计提案)
    ↓ (被接受)
ADR Document (最终决策)
    ↓ (开始实施)
GitHub Projects (任务追踪)
    ↓ (创建 Issues)
Implementation (代码实现)
```

---

## 📌 交叉引用最佳实践

在所有文档中维护交叉引用以确保可追踪性：

**在 RFC 中：**
```markdown
## Related
- Discussion: https://github.com/xukin10/AetherOS/discussions/XXX
- Linked Issue: #XXX
```

**在 ADR 中：**
```markdown
## Related Decisions
- RFC: docs/rfcs/RFC-001-example.md
- Previous ADR: docs/adrs/ADR-XXX-example.md
```

**在 Issues 中：**
```markdown
Related ADR: docs/adrs/ADR-001-example.md
Related RFC: #XXX (PR)
```

---

## 📊 文件位置

```
AetherOS/
├── docs/
│   ├── rfcs/                    # RFC 提案
│   │   ├── RFC-TEMPLATE.md
│   │   └── RFC-001-example.md
│   ├── adrs/                    # ADR 决策记录
│   │   ├── ADR-TEMPLATE.md
│   │   ├── ADR-001-example.md
│   │   └── README.md            # ADR 索引
│   └── GOVERNANCE.md            # 本文件
├── .github/
│   ├── discussions/
│   │   └── welcome.md           # Discussions 欢迎文本
│   └── ISSUE_TEMPLATE/          # Issue 模板
└── PROJECT.md                   # 项目概览
```

---

## 🏷️ GitHub 标签约定

建议创建以下标签以提高组织性：

| 标签 | 说明 |
|------|------|
| `type:rfc` | RFC 提案 |
| `type:adr` | ADR 相关 |
| `type:feature` | 新功能 |
| `type:bug` | 缺陷修复 |
| `status:needs-review` | 需要审查 |
| `status:accepted` | 已接受 |
| `status:rejected` | 已拒绝 |
| `priority:high` | 高优先级 |
| `priority:low` | 低优先级 |

---

## 💡 示例工作流程

### 场景：提议新的插件系统架构

1. **[Week 1] Discussion 阶段**
   - 在 Discussions 中发起："Should we adopt a plugin system?"
   - 收集团队和社区反馈

2. **[Week 2-3] RFC 阶段**
   - 创建 `RFC-002-plugin-architecture.md`
   - 提交 PR，标记为 `type:rfc`
   - 迭代评论，完善设计

3. **[Week 4] ADR 阶段**
   - 创建 `ADR-002-plugin-architecture.md`
   - 标记 RFC PR 为 `status:accepted`
   - 合并 ADR 到 main

4. **[Week 5+] Projects 阶段**
   - 创建 Project 卡片
   - 拆分为多个 Issue：
     - Issue #10: 设计 API
     - Issue #11: 实现加载机制
     - Issue #12: 编写文档
   - 追踪进度

---

## ✅ 检查清单

启动新项目时：

- [ ] 启用 GitHub Discussions
- [ ] 创建 `/docs/rfcs/` 和 `/docs/adrs/` 目录
- [ ] 添加 RFC 和 ADR 模板
- [ ] 创建 GitHub Project 看板
- [ ] 定义标签系统
- [ ] 发布治理指南文档
- [ ] 举办团队培训会议

---

## 📚 参考资源

- [ADR 介绍](https://adr.github.io/)
- [RFC 流程示例](https://en.wikipedia.org/wiki/Request_for_Comments)
- [GitHub Discussions](https://docs.github.com/en/discussions)
- [GitHub Projects](https://docs.github.com/en/issues/planning-and-tracking-with-projects)

---

**最后更新：** 2026-07-01  
**维护者：** @xukin10
