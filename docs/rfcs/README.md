# Request for Comments (RFCs)

本目录包含 AetherOS 项目的所有 RFC 提案。

## 什么是 RFC？

RFC（Request for Comments）是在实施前展示和讨论重要设计的正式过程。RFC 提案通过 Pull Request 发布，允许社区提供反馈。

## 📋 RFC 索引

| # | 标题 | 状态 | PR | 日期 |
|---|------|------|----|----|
| 001 | [示例：新功能提案](./RFC-001-example.md) | Proposed | #XXX | TBD |

## 状态说明

- **Proposed** - 初次提案，等待反馈
- **Under Review** - 正在审查，有积极讨论
- **Accepted** - 被接受，将创建对应 ADR 并开始实施
- **Rejected** - 被拒绝，不会实施
- **Deprecated** - 之前被接受，但现在不再推荐

## RFC 生命周期

```
Proposed → Under Review → Accepted → Implementation (via ADR) → Done
                       ↘ Rejected (Archive)
```

## 创建新 RFC

### 步骤 1: 讨论阶段
在 GitHub Discussions 中提起你的想法，收集初步反馈。

### 步骤 2: 创建 RFC 文档
1. 复制 `/docs/rfcs/RFC-TEMPLATE.md`
2. 命名为 `RFC-XXX-[title].md`（按顺序递增）
3. 填充完整的设计文档
4. 包括所有替代方案和权衡

### 步骤 3: 提交 Pull Request
```bash
# 创建 RFC 分支
git checkout -b rfc/001-new-feature

# 提交 RFC 文件
git add docs/rfcs/RFC-001-new-feature.md
git commit -m "RFC: Propose new feature"

# 推送并创建 PR
git push origin rfc/001-new-feature
```

PR 描述应包含：
- 相关 Discussion 链接（如有）
- RFC 的简要总结
- 关键决策点
- 需要社区反馈的地方

标记 PR 为 `type:rfc`

### 步骤 4: 收集反馈
- 回应所有评论
- 根据反馈迭代设计
- 更新 RFC 文档

### 步骤 5: 最终决策
核心团队进行最后审查并做出决策：
- ✅ **接受** - 标记 PR 为 `status:accepted`，创建对应的 ADR
- ❌ **拒绝** - 标记 PR 为 `status:rejected`，保存在历史记录中

## 查看相关文件

- **ADR 记录** - `/docs/adrs/`
- **治理指南** - `/docs/GOVERNANCE.md`
- **讨论** - GitHub Discussions tab

## 最佳实践

✅ **DO:**
- 清晰地阐述动机和问题
- 考虑多个设计选项
- 讨论权衡和后果
- 寻求社区反馈
- 保持 RFC 专注于单一决策

❌ **DON'T:**
- 在 RFC 中混入多个不相关的决策
- 跳过替代方案讨论
- 忽视社区反馈
- 仓促做出决策

## 参考

- [RFC Wikipedia](https://en.wikipedia.org/wiki/Request_for_Comments)
- [Rust RFC Process](https://rust-lang.github.io/rfcs/)
- [GitHub Discussions Best Practices](https://docs.github.com/en/discussions/guides/best-practices-for-community-conversations-on-github)

---

**最后更新：** 2026-07-01
