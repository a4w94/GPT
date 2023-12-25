# gpt

`gpt` is a command-line tool that provides various functionalities for code-related tasks. It is built using Cobra, a popular Go library for creating powerful modern CLI applications.

## Features

- **Summary:** Provides the article summary.
- **Refactor:** Assists in refactoring code.
- **Translate:** Translates the input paragraph.

## Installation

Before using `gpt`, make sure to set up your environment variables by creating a `.env` file in the project root and populating it with the required values.

```sh
# .env file
API_KEY=your_api_key
```

## Usage

### Translate

```sh
 go run main.go plugin -n translate
```
Use this command to send a translate request to GPT for the provided paragraph.

### Summary

```sh
 go run main.go plugin -n summary
```
Use this command to send a summary request to GPT for the provided article.

### Refactor
```sh
./gpt refactor
```
or
```sh
go run main.go refactor
```

Use this command to send a refactoring request to GPT for the provided code.


### 自行新增

```markdown
- 步驟 1: 複製目錄
- 步驟 2: 更改名稱為您要的功能名稱並確保目錄格式

    plugins/
    └── {function_name}/
        ├── config.json
        ├── {function_name}.json
        └── {function_name}.tmpl

- 步驟 3: 調整config.json參數
- 步驟 4: 依照您的需求修改{function_name}.tmpl並確保{function_name}.json包含{function_name}.tmpl所有key值
- 步驟 5: 執行 `go run main.go plugin -n {function_name}`
```

# Demo

## 1.
```php

    /**
     * 取得活動資料與活動設定
     *
     * @param int $hallId 廳主ID
     * @param int $activityId 活動ID
     * @return array|bool|string
     */
    public function getActivityDetail($hallId, $activityId)
    {
        // 撈取活動資料
        $originalActivityInfo = $this->getActivityInfo($hallId, $activityId);
        if ($originalActivityInfo === false) {
            return false;
        } elseif ($originalActivityInfo['activity_type'] != 3) {
            // 活動類型不符
            return 'activityTypeError';
        }

        // 撈取活動設定
        $originalActivitySetting = $this->activityRepositoryDepositbet->getActivitySetting($activityId);

        $taskSetting = json_decode($originalActivityInfo['task_setting'], true);
        $mission = [];
        foreach ($taskSetting as $taskVal) {
            $mission[$taskVal['sort']] = [
                'name' => $taskVal['name'], // 任務名稱
                'depositUserLimit' => ($taskVal['deposit_user_limit'] != 0) ? $taskVal['deposit_user_limit'] : '', // 派發人數限制
            ];
        }

        // 取得廳主開放幣別列表
        $hallCurrencyList = $this->coreCurrency->getCurrencyList($hallId);

        foreach ($originalActivitySetting as $setVal) {
            // 跳過已經不存在的幣別
            if (!collect($hallCurrencyList)->has($setVal['currency'])) {
                continue;
            }

            // 人民幣除顯示(RMB)外，其他統一為CNY
            $currency = ($setVal['currency'] == 'CNY') ? 'RMB' : $setVal['currency'];

            $rewardList = json_decode($setVal['reward_list'], true);
            $rewardSetting = [];
            $complexType = '';
            foreach ($rewardList as $rewardVal) {
                $rewardSetting[] = [
                    'totalDeposit' => ($rewardVal['total_deposit'] != 0) ? $rewardVal['total_deposit'] : '', // 累計存款金額
                    'totalBet' => ($rewardVal['total_bet'] != 0) ? $rewardVal['total_bet'] : '', // 累計打碼
                    'amount' => ($rewardVal['amount'] != 0) ? $rewardVal['amount'] : '', // 可領獎勵
                    'percent' => ($rewardVal['percent'] != 0) ? $rewardVal['percent'] : '', // 獎勵比例
                    'amountLimit' => ($rewardVal['amount_limit'] != 0) ? $rewardVal['amount_limit'] : '', // 獎勵上限
                    'complexPercent' => ($rewardVal['complex_percent'] != 0) ? $rewardVal['complex_percent'] : (($originalActivityInfo['audit_type'] === 2) ? 0 : ''), // 打碼倍數，如果「未設定」且稽核設定是「照個任務設定」必須顯示為0
                    'complexAmount' => ($rewardVal['complex_amount'] != 0) ? $rewardVal['complex_amount'] : (($originalActivityInfo['audit_type'] === 2) ? 0 : ''), // 打碼量，如果「未設定」且稽核設定是「照個任務設定」必須顯示為0
                    'totalPositiveProfit' => (isset($rewardVal['total_positive_profit']) && $rewardVal['total_positive_profit'] != 0) ? $rewardVal['total_positive_profit'] : '', // 正盈利金額
                    'totalNegativeProfit' => (isset($rewardVal['total_negative_profit']) && $rewardVal['total_negative_profit'] != 0) ? $rewardVal['total_negative_profit'] : '', // 負盈利金額
                    'numberDeposit' => (isset($rewardVal['number_deposit']) && $rewardVal['number_deposit'] != 0) ? $rewardVal['number_deposit'] : '', // 存款次數
                    'totalWithdraw' => (isset($rewardVal['total_withdraw']) && $rewardVal['total_withdraw'] != 0) ? $rewardVal['total_withdraw'] : '', // 累計取款
                    'numberWithdraw' => (isset($rewardVal['number_withdraw']) && $rewardVal['number_withdraw'] != 0) ? $rewardVal['number_withdraw'] : '', // 取款次數
                ];

                // 判斷打碼稽核種類
                if ($originalActivityInfo['audit_type'] == 2) {
                    // 當稽核設定為「照各任務設定」才可設定任務裡面的打碼稽核種類
                    $complexType = $setVal['complex_type'] ?? '';
                }
            }

            $mission[$setVal['sort']]['detail'][$currency] = [
                'depositCategoryList' => json_decode($setVal['deposit_category_list'], true), // 指定存款類別
                'betGameList' => json_decode($setVal['bet_game_list'], true), // 指定打碼產品
                'conditionSetting' => json_decode($setVal['condition_setting'], true), // 條件設置Key(deposit:存款, bet:打碼, positive_profit:正盈利, negative_profit:負盈利)
                'betConditionType' => $setVal['bet_condition_type'], // 領獎打碼條件(1:累計打碼 2:依會員存款倍數(固定) 3:依會員取款倍數(固定))
                'betConditionMultiple' => $setVal['bet_condition_multiple'], // 會員存取款倍數
                'rewardSystem' => $setVal['reward_system'], // 獎勵制度(1:累計制, 2:僅領取最高)
                'rewardType' => $setVal['reward_type'], // 獎勵方式(1:單筆彩金, 2:存款比例, 3:打碼比例, 4:正盈利比例, 5:負盈利比例, 6:取款比例)
                'rewardSetting' => $rewardSetting, // 獎勵設定
                'complexType' => $complexType, // 打碼稽核種類(1:元, 2:倍)
                'withdrawCategoryList' => json_decode($setVal['withdraw_category_list'], true), // 指定取款類別
                'gameWhiteListTagId' => $setVal['game_whitelist_tag_id'], // 指定遊戲名單群組id
            ];
        }

        $startDate = date('Y-m-d', strtotime($originalActivityInfo['start_time']));

        $endDate = '';
        if (isset($originalActivityInfo['end_time'])) {
            $endDate = date('Y-m-d', strtotime($originalActivityInfo['end_time']));
        }

        // 取得參與打碼要求設定
        switch ($originalActivityInfo['bet_type']) {
            case 2:
                // 固定打碼金額
                $betMinimum = ($originalActivityInfo['bet_minimum'] != 0) ? $originalActivityInfo['bet_minimum'] : '';
                break;
            default:
                // 依存款設定倍數、依總存設定倍數
                $betMinimum = ($originalActivityInfo['bet_multiple'] != 0) ? $originalActivityInfo['bet_multiple'] : '';
                break;
        }

        $targetIdList = [];
        $targetMemberList = [];
        if ($originalActivityInfo['target_type'] == 3) {
            //3.指定會員
            $targetMemberListDB = $this->activityRepositoryDepositbet->getTargetUser($originalActivityInfo['id']);
            foreach ($targetMemberListDB as $val) {
                $targetMemberList[] = $val['username'];
            }
        } else {
            //1:層級, 2:等級
            $targetIdList = json_decode($originalActivityInfo['target_id_list'], true);
        }

        // 等級類別種類(general:終身等級 shortTerm:短期等級)
        $categoryType = '';
        $categoryExist = '';
        if ($originalActivityInfo['category'] != 0) {
            $categoryDetail = $this->coreVIPCategory->getCategoryDetail($hallId, $originalActivityInfo['category']);
            $categoryExist = 'N';
            if (isset($categoryDetail['isShortTerm'])) {
                $categoryType = $this->coreVIPCategory->getCategoryType($categoryDetail['isShortTerm']);
                $categoryExist = 'Y';
            }
        }

        $activityInfo = [
            'id' => $originalActivityInfo['id'], // 活動ID
            'name' => json_decode($originalActivityInfo['name'], true), // 活動名稱
            'oneOffSwitch' => $originalActivityInfo['one_off_switch'], // 一次性活動(Y/N)
            'activityType' => $originalActivityInfo['activity_type'], // 活動類型
            'startDate' => $startDate, // 活動開始時間
            'endDate' => $endDate, // 活動結束時間
            'cycleSetting' => $originalActivityInfo['settlement_type'], // 循環設定(0:無循環, 1:每日, 2:每週, 3:雙週, 4:每月, 5:指定日期)
            'cycleTime' => $originalActivityInfo['settlement_detail'] ?? '', // 循環日期
            'receivedType' => $originalActivityInfo['received_type'], // 領取方式(S:系統派發, R:玩家自領)
            'isAutoPayment' => ($originalActivityInfo['received_type'] == 'R') ? $originalActivityInfo['is_auto_payment'] : '', // 自領到期後是否自動派發(Y/N)
            'finalReceivedDay' => ($originalActivityInfo['received_type'] == 'R') ? (($originalActivityInfo['final_received_day'] != 0) ? $originalActivityInfo['final_received_day'] : '') : '', // 派發後幾天截止
            'ipLimitCount' => ($originalActivityInfo['received_type'] == 'R') ? (($originalActivityInfo['ip_limit_count'] != 0) ? $originalActivityInfo['ip_limit_count'] : '') : '', // IP限制次數
            'realnameLimitCount' => ($originalActivityInfo['realname_limit_count'] != 0) ? $originalActivityInfo['realname_limit_count'] : '', // 真實姓名限制次數
            'auditType' => $originalActivityInfo['audit_type'], // 稽核設定(0:無稽核, 1:固定倍率, 2:照各任務設定)
            'complex' => ($originalActivityInfo['complex'] != 0) ? $originalActivityInfo['complex'] : '', // 打碼倍數
            'commissionSwitch' => $originalActivityInfo['commission_switch'], // 是否寫入退佣(Y/N)
            'requirement' => ($originalActivityInfo['is_application'] == 'Y') ? 'R' : 'S', // 報名方式(S:不需報名, R:玩家申請)
            'reviewType' => ($originalActivityInfo['is_manual_review'] == 'Y') ? 'M' : 'S', // 審核方式(S:自動審核, M:人工審核)
            'reviewDay' => ($originalActivityInfo['review_day'] != 0) ? $originalActivityInfo['review_day'] : '', // 審核天數
            'depositMinimum' => ($originalActivityInfo['deposit_minimum'] != 0) ? $originalActivityInfo['deposit_minimum'] : '', // 參與存款要求(人民幣)
            'betType' => $originalActivityInfo['bet_type'], // 參與打碼要求種類(1:依存款設定倍數, 2:固定打碼金額, 3:依總存設定倍數)
            'betMinimum' => $betMinimum, // betType:2 => 參與存款要求(人民幣), betType:3 => 會員總存倍數
            'betGameList' => json_decode($originalActivityInfo['bet_game_list'], true), // 打碼指定產品
            'finishActivityList' => json_decode($originalActivityInfo['finish_activity_list'], true), // 完成優惠活動ID列表
            'excludeActivityList' => json_decode($originalActivityInfo['exclude_activity_list'], true), // 排除優惠活動ID列表
            'mission' => array_values($mission), // 任務
            'targetType' => $originalActivityInfo['target_type'], // 派發對象(1:層級, 2:等級, 3:指定會員)
            'vipCategoryId' => ($originalActivityInfo['category'] != 0) ? $originalActivityInfo['category'] : '', // 等級類別(1:累計打碼)
            'vipCategoryExist' => $categoryExist, // 等級類別是否存在
            'categoryType' => $categoryType, // 等級類別種類(general:終身等級 shortTerm:短期等級)
            'targetIdList' => $targetIdList, // 派發對象列表
            'targetMemberList' => $targetMemberList, // 指定會員列表
            'status' => $originalActivityInfo['status'], // 活動狀態(1:未開始, 2:進行中, 3:已停止)
            'originalActivityId' => $originalActivityInfo['original_activity_id'] ?? 0, // 原始活動ID (0 表示此活動未曾在活動開始後進行修改)
            'blacklistSwitch' => $originalActivityInfo['blacklist_switch'], // 是否排除黑名單
            'gameBlackListTagId' => $originalActivityInfo['game_blacklist_tag_id'], // 遊戲黑名單群組id
            'gameWhiteListTagId' => $originalActivityInfo['game_whitelist_tag_id'], // 指定遊戲名單群組id
        ];

        return $activityInfo;
    }
```
## 2. Demo CodeGPT
### 安裝
```bash
brew install codegpt
```
or
```bash
go install github.com/appleboy/CodeGPT/cmd/codegpt@latest
```
### 設定api_key
```bash
codegpt config set openai.api_key sk-xxxxxxx
```
### 設定model (不一定要)
```bash
codegpt config set openai.model gpt-3.5-turbo-16k
```
### 執行
```bash
git add .
```

請gpt幫忙寫commit
```bash
Codegpt commit –preview //
```

請gpt幫忙code review 並翻譯
```bash
codegpt commit --lang zh-tw –preview
```

提交剛剛的commit
```bash
codegpt commit –amend
```

使用自訂格式commit
```bash
codegpt commit --preview --template_file ./commit_message.tpl --template_vars_file ./commit_vars.env
```







