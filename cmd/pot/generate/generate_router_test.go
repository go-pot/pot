package generate

import (
	"encoding/json"
	"testing"
)

func TestGenerateRouter(t *testing.T) {

	var s interface{}
	json.Unmarshal([]byte(d), &s)
	GenerateRouter("", s)
}

var d = `{
    "swagger": "2.0",
    "info": {
        "title": "wjs api",
        "description": "所有金额单位为 \"分\" 如 1元(100分), 时间格式为 \"RFC3339\" 如2006-01-02T15:04:05Z07:00, 区间范围 包括左边界不包括右边界",
        "version": "1.0.0",
        "contact": {}
    },
    "basePath": "/v2",
    "paths": {
        "/balance/chargesubmit": {
            "post": {
                "tags": [
                    "/balance"
                ],
                "summary": "充值申请",
                "description": "充值申请",
                "operationId": "/balance/chargesubmit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eMobile: 手机号 \u003cbr/\u003eBankcard: 银行卡号 \u003cbr/\u003eNumber: 金额 \u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqChargeRecord"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/balance/getbalancechangerecord": {
            "post": {
                "tags": [
                    "/balance"
                ],
                "summary": "资金变更记录",
                "description": "资金变更记录",
                "operationId": "/balance/getbalancechangerecord",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBeginTime: 开始时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eOffset: 偏移 \u003cbr/\u003eLimit: 限制 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserTimeRangeOffsetLimit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eAccount: 账号 \u003cbr/\u003eUserName: 实名 \u003cbr/\u003eNumber: 变更金额 \u003cbr/\u003eMsg: 交易信息 \u003cbr/\u003eCreateTime: 时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespUserBalanceChange"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/balance/getbalancechangerecordbyquery": {
            "post": {
                "tags": [
                    "/balance"
                ],
                "summary": "资金变更记录",
                "description": "资金变更记录",
                "operationId": "/balance/getbalancechangerecordbyquery",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eFromType: 类型 \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetBalanceRecord"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eAccount: 账号 \u003cbr/\u003eUserName: 实名 \u003cbr/\u003eNumber: 变更金额 \u003cbr/\u003eAfterNumber: 变更后的金额 \u003cbr/\u003eMsg: 交易信息 \u003cbr/\u003eCreateTime: 时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespUserBalanceChangeRecord"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/balance/getuserassets": {
            "post": {
                "tags": [
                    "/balance"
                ],
                "summary": "GetUserAssets",
                "description": "购买记录",
                "operationId": "/balance/getuserassets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eUserName: 实名 \u003cbr/\u003eAccount: 账号 \u003cbr/\u003eBeginTime: 开始时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eOffset: 偏移 \u003cbr/\u003eLimit: 限制 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserTime"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eAccount: 账号 \u003cbr/\u003eUserName: 实名 \u003cbr/\u003eNumber: 可用余额 \u003cbr/\u003eFreezeNumber: 冻结金额 \u003cbr/\u003eTotalNumber: 总投资金额 \u003cbr/\u003eTotalPayNumber: 总实付金额 \u003cbr/\u003eTotalReturnNumber: 总回款本金 \u003cbr/\u003eTotalEarningsNumber: 总回款收益 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespUserAssets"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/balance/getuserassetscountnew": {
            "post": {
                "tags": [
                    "/balance"
                ],
                "summary": "GetUserAssetsCountNew",
                "description": "购买记录总数 2017-11-12",
                "operationId": "/balance/getuserassetscountnew",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBeginTime: 开始时间 \u003cbr/\u003eEndTime: 结束时间 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserTimeRange"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/balance/getuserassetsnew": {
            "post": {
                "tags": [
                    "/balance"
                ],
                "summary": "GetUserAssetsNew",
                "description": "购买记录 2017-11-12",
                "operationId": "/balance/getuserassetsnew",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBeginTime: 开始时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eOffset: 偏移 \u003cbr/\u003eLimit: 限制 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserTimeRangeOffsetLimit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eAccount: 账号 \u003cbr/\u003eUserName: 实名 \u003cbr/\u003eNumber: 可用余额 \u003cbr/\u003eFreezeNumber: 冻结金额 \u003cbr/\u003eTotalNumber: 总投资金额 \u003cbr/\u003eTotalPayNumber: 总实付金额 \u003cbr/\u003eTotalReturnNumber: 总回款本金 \u003cbr/\u003eTotalEarningsNumber: 总回款收益 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespUserAssets"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/balance/withdraw": {
            "post": {
                "tags": [
                    "/balance"
                ],
                "summary": "提现",
                "description": "扣除余额,放入冻结金额,新增提现记录",
                "operationId": "/balance/withdraw",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBankcardId: 银行卡id \u003cbr/\u003eNumber: 金额 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqWithdraw"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBankcardId: 银行卡id \u003cbr/\u003eNumber: 金额 ",
                        "schema": {
                            "$ref": "#/definitions/models.ReqWithdraw"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/bankcard/addcard": {
            "post": {
                "tags": [
                    "/bankcard"
                ],
                "summary": "添加银行卡",
                "description": "添加银行卡，已存在会添加失败",
                "operationId": "/bankcard/addcard",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eCardNumber: \u003cbr/\u003eBankCode: 银行编码 \u003cbr/\u003eBankName: 银行名称 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqAddCard"
                        }
                    }
                ]
            }
        },
        "/bankcard/getbankcard": {
            "post": {
                "tags": [
                    "/bankcard"
                ],
                "summary": "获取银行卡根据id",
                "description": "获取银行卡根据id",
                "operationId": "/bankcard/getbankcard",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"银行卡id\"\u003cbr/\u003eBankcardId: 银行卡id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBankcardId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "银行卡信息\u003cbr/\u003eId: 银行卡id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eState: 状态 \u003cbr/\u003ePayWay: 支付方式 \u003cbr/\u003eTotalIn: 最大流入限额 \u003cbr/\u003eTotalOut: 最大流出限额 \u003cbr/\u003ePhone: 预留手机号 \u003cbr/\u003eCard: 银行卡号 \u003cbr/\u003eBankCode: 银行代号 \u003cbr/\u003eBindId: 绑定id \u003cbr/\u003eBankName: 银行名 \u003cbr/\u003eBankDesc: 银行描述 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.Bankcard"
                        }
                    },
                    "202": {
                        "description": "查询失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/bankcard/getbankcardbycard": {
            "post": {
                "tags": [
                    "/bankcard"
                ],
                "summary": "获取银行卡根据卡号",
                "description": "获取银行卡根据卡号",
                "operationId": "/bankcard/getbankcardbycard",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"银行卡id\"\u003cbr/\u003eCard: 银行卡号 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBankcardCard"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "银行卡信息\u003cbr/\u003eId: \u003cbr/\u003eUserId: \u003cbr/\u003eState: \u003cbr/\u003ePayWay: \u003cbr/\u003eTotalIn: \u003cbr/\u003eTotalOut: \u003cbr/\u003ePhone: \u003cbr/\u003eCard: \u003cbr/\u003eBankCode: \u003cbr/\u003eBindId: \u003cbr/\u003eBankName: \u003cbr/\u003eBankDesc: \u003cbr/\u003eCreateTime: \u003cbr/\u003eSingleLimit: \u003cbr/\u003eDayLimit: ",
                        "schema": {
                            "$ref": "#/definitions/models.RespBankcard"
                        }
                    },
                    "202": {
                        "description": "查询失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/bankcard/getcardlimit": {
            "post": {
                "tags": [
                    "/bankcard"
                ],
                "summary": "获取支付限额",
                "description": "获取支付限额",
                "operationId": "/bankcard/getcardlimit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"银行编码,支付渠道\"\u003cbr/\u003ePayWay: 渠道 \u003cbr/\u003eBankCode: 银行编码 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqCardInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "支付限额\u003cbr/\u003eId: 银行卡限制id \u003cbr/\u003ePayWay: 支付方式 \u003cbr/\u003eBankCode: 银行代号 \u003cbr/\u003eBankName: 银行名字 \u003cbr/\u003eSingleLimit: 单笔限额 \u003cbr/\u003eDayLimit: 天数 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.CardLimit"
                        }
                    },
                    "202": {
                        "description": "查询失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/bankcard/getcardsupport": {
            "post": {
                "tags": [
                    "/bankcard"
                ],
                "summary": "获取支持的银行卡",
                "description": "获取支持的银行卡",
                "operationId": "/bankcard/getcardsupport",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "支付限额\u003cbr/\u003eId: 银行卡限制id \u003cbr/\u003ePayWay: 支付方式 \u003cbr/\u003eBankCode: 银行代号 \u003cbr/\u003eBankName: 银行名字 \u003cbr/\u003eSingleLimit: 单笔限额 \u003cbr/\u003eDayLimit: 天数 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.CardLimit"
                        }
                    },
                    "202": {
                        "description": "查询失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/bankcard/updatestate": {
            "post": {
                "tags": [
                    "/bankcard"
                ],
                "summary": "修改绑卡状态",
                "description": "修改绑卡状态",
                "operationId": "/bankcard/updatestate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"银行编码,支付渠道\"\u003cbr/\u003eBankcardId: 银行卡id \u003cbr/\u003eState: 银行卡号 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBankcardState"
                        }
                    }
                ]
            }
        },
        "/banner/add": {
            "post": {
                "tags": [
                    "/banner"
                ],
                "summary": "添加banner",
                "description": "添加banner",
                "operationId": "/banner/add",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eState: 轮播图状态号 \u003cbr/\u003ePlatform: 平台id \u003cbr/\u003ePic: 图片url \u003cbr/\u003eUrl: 跳转url ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBannerAdd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/banner/countallbanner": {
            "post": {
                "tags": [
                    "/banner"
                ],
                "summary": "所有banner列表",
                "description": "所有banner列表",
                "operationId": "/banner/countallbanner",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/banner/getallbanner": {
            "post": {
                "tags": [
                    "/banner"
                ],
                "summary": "所有banner列表",
                "description": "所有banner列表",
                "operationId": "/banner/getallbanner",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eState: 状态 \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBannerList"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eId: 轮播id \u003cbr/\u003eState: 状态 \u003cbr/\u003ePlatform: 平台id \u003cbr/\u003ePic: 图片地址 \u003cbr/\u003eUrl: 点击跳转地址 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.Banner"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/banner/list": {
            "post": {
                "tags": [
                    "/banner"
                ],
                "summary": "获取 APP banner列表",
                "description": "获取 APP banner列表",
                "operationId": "/banner/list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eId: 轮播id \u003cbr/\u003eState: 状态 \u003cbr/\u003ePlatform: 平台id \u003cbr/\u003ePic: 图片地址 \u003cbr/\u003eUrl: 点击跳转地址 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.Banner"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/banner/listupdate": {
            "post": {
                "tags": [
                    "/banner"
                ],
                "summary": "修改banner显示列表",
                "operationId": "/banner/listupdate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/banner/sublist": {
            "post": {
                "tags": [
                    "/banner"
                ],
                "summary": "获取 APP 副banner列表",
                "description": "获取 APP 副banner列表",
                "operationId": "/banner/sublist",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eId: 轮播id \u003cbr/\u003eState: 状态 \u003cbr/\u003ePlatform: 平台id \u003cbr/\u003ePic: 图片地址 \u003cbr/\u003eUrl: 点击跳转地址 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.Banner"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/banner/update": {
            "post": {
                "tags": [
                    "/banner"
                ],
                "summary": "修改banner",
                "description": "修改banner",
                "operationId": "/banner/update",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eId: 轮播图id \u003cbr/\u003eState: 轮播图状态号 \u003cbr/\u003ePlatform: 平台id \u003cbr/\u003ePic: 图片url \u003cbr/\u003eUrl: 跳转url ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBannerUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/bids/addbid": {
            "post": {
                "tags": [
                    "/bids"
                ],
                "summary": "添加标的物 立即起息",
                "description": "添加标的物 立即起息",
                "operationId": "/bids/addbid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"添加标\"\u003cbr/\u003eBidMark: 标的购买限制 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eBidNo: 标订单号 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标的详情地址 \u003cbr/\u003eBidRate: 标的利率 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eCheckDate: 结算日期 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eMinNumber: 起投金额 \u003cbr/\u003eMaxNumber: 最大金额 \u003cbr/\u003eStepNumber: 增幅金额 \u003cbr/\u003eMaxBuySize: 最大购买笔数 \u003cbr/\u003eRefillDaysLimit: 标的物最大期限 \u003cbr/\u003eRefillBidType: 标计息类型 \u003cbr/\u003eRefillBidMark: 标的购买限制 \u003cbr/\u003eRefillRemainNumberMin: 全部符合条件的标 累计小于最小续标金额 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqAddBid"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "标的id\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/bids/addbidnew": {
            "post": {
                "tags": [
                    "/bids"
                ],
                "summary": "添加标 2017-08-25",
                "description": "添加标",
                "operationId": "/bids/addbidnew",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"添加标\"\u003cbr/\u003eBidType: 标的类型 \u003cbr/\u003eBidMark: 标的购买限制 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eBidNo: 标订单号 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标的详情地址 \u003cbr/\u003eBidRate: 标的利率 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eMinNumber: 起投金额 \u003cbr/\u003eMaxNumber: 最大金额 \u003cbr/\u003eStepNumber: 增幅金额 \u003cbr/\u003eMaxBuySize: 最大购买笔数 \u003cbr/\u003eRefillDaysLimit: 标的物最大期限 \u003cbr/\u003eRefillBidType: 标计息类型 \u003cbr/\u003eRefillBidMark: 标的购买限制 \u003cbr/\u003eRefillRemainNumberMin: 全部符合条件的标 累计小于最小续标金额 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqAddBidNew"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "标的id\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/bids/addbidpool": {
            "post": {
                "tags": [
                    "/bids"
                ],
                "summary": "添加标的物池",
                "description": "添加标的物池",
                "operationId": "/bids/addbidpool",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"添加活期标\"\u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eBidNo: 标订单号 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eBidInfoPath: 标的详情地址 \u003cbr/\u003eCheckDate: 结算日期 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqAddBidPool"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "标的id\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/bids/addbidsatisfy": {
            "post": {
                "tags": [
                    "/bids"
                ],
                "summary": "添加标的物 满标起息",
                "description": "添加标的物 满标起息",
                "operationId": "/bids/addbidsatisfy",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"添加标\"\u003cbr/\u003eBidMark: 标的购买限制 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eBidNo: 标订单号 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标的详情地址 \u003cbr/\u003eBidRate: 标的利率 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eMinNumber: 起投金额 \u003cbr/\u003eMaxNumber: 最大金额 \u003cbr/\u003eStepNumber: 增幅金额 \u003cbr/\u003eMaxBuySize: 最大购买笔数 \u003cbr/\u003eRefillDaysLimit: 标的物最大期限 \u003cbr/\u003eRefillBidType: 标计息类型 \u003cbr/\u003eRefillBidMark: 标的购买限制 \u003cbr/\u003eRefillRemainNumberMin: 全部符合条件的标 累计小于最小续标金额 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqAddBidSatisfy"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "标的id\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/bids/getbid": {
            "post": {
                "tags": [
                    "/bids"
                ],
                "summary": "获取标的根据标的id",
                "description": "获取标的根据标的id",
                "operationId": "/bids/getbid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"标id\"\u003cbr/\u003eBidId: 标的id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBidId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "标数据\u003cbr/\u003eId: 标的id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003eBidType: 标计息类型 标记活期定期 \u003cbr/\u003eBidMark: 标的购买限制 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eRemainNumber: 标剩余金额 \u003cbr/\u003eBidNo: 标订单号 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标详情路径 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eCheckDate: 结算日期 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 \u003cbr/\u003eIsVip: 是计入vip的 \u003cbr/\u003eIsRank: 是计入rank的 \u003cbr/\u003eMark: 标记 \u003cbr/\u003eDays: 实际经过的天数 \u003cbr/\u003eBidRate: 标的物利率 \u003cbr/\u003eIncomeDate: 起息时间 \u003cbr/\u003eBidRateBase: 基础利率 \u003cbr/\u003eBidRateRaise: 加息利率 \u003cbr/\u003eMinNumber: 起投金额 \u003cbr/\u003eMaxNumber: 最大金额 \u003cbr/\u003eStepNumber: 增幅金额 \u003cbr/\u003eMaxBuySize: 最大购买笔数 \u003cbr/\u003eRefillDaysLimit: 标的物最大期限 \u003cbr/\u003eRefillBidType: 标计息类型 \u003cbr/\u003eRefillBidMark: 标的购买限制 \u003cbr/\u003eRefillRemainNumberMin: 全部符合条件的标 累计小于最小续标金额 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespBids"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/bids/getbids": {
            "post": {
                "tags": [
                    "/bids"
                ],
                "summary": "获取可以显示未满标的标的物",
                "description": "获取可以显示未满标的标的物",
                "operationId": "/bids/getbids",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "标数据数组\u003cbr/\u003eId: 标的id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003eBidType: 标计息类型 标记活期定期 \u003cbr/\u003eBidMark: 标的购买限制 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eRemainNumber: 标剩余金额 \u003cbr/\u003eBidNo: 标订单号 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标详情路径 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eCheckDate: 结算日期 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 \u003cbr/\u003eIsVip: 是计入vip的 \u003cbr/\u003eIsRank: 是计入rank的 \u003cbr/\u003eMark: 标记 \u003cbr/\u003eDays: 实际经过的天数 \u003cbr/\u003eBidRate: 标的物利率 \u003cbr/\u003eIncomeDate: 起息时间 \u003cbr/\u003eBidRateBase: 基础利率 \u003cbr/\u003eBidRateRaise: 加息利率 \u003cbr/\u003eMinNumber: 起投金额 \u003cbr/\u003eMaxNumber: 最大金额 \u003cbr/\u003eStepNumber: 增幅金额 \u003cbr/\u003eMaxBuySize: 最大购买笔数 \u003cbr/\u003eRefillDaysLimit: 标的物最大期限 \u003cbr/\u003eRefillBidType: 标计息类型 \u003cbr/\u003eRefillBidMark: 标的购买限制 \u003cbr/\u003eRefillRemainNumberMin: 全部符合条件的标 累计小于最小续标金额 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespBids"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/bids/getbidsbyquery": {
            "post": {
                "tags": [
                    "/bids"
                ],
                "summary": "获取全部标的物",
                "description": "获取全部标的物",
                "operationId": "/bids/getbidsbyquery",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"获取标的物\"\u003cbr/\u003eQuery: 搜索 \u003cbr/\u003eBidType: 标的类型 \u003cbr/\u003eBidMark: 标的分类 \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetBidsByQuery"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "标数据数组\u003cbr/\u003eId: 标的id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003eBidType: 标计息类型 标记活期定期 \u003cbr/\u003eBidMark: 标的购买限制 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eRemainNumber: 标剩余金额 \u003cbr/\u003eBidNo: 标订单号 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标详情路径 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eCheckDate: 结算日期 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 \u003cbr/\u003eIsVip: 是计入vip的 \u003cbr/\u003eIsRank: 是计入rank的 \u003cbr/\u003eMark: 标记 \u003cbr/\u003eDays: 实际经过的天数 \u003cbr/\u003eBidRate: 标的物利率 \u003cbr/\u003eIncomeDate: 起息时间 \u003cbr/\u003eBidRateBase: 基础利率 \u003cbr/\u003eBidRateRaise: 加息利率 \u003cbr/\u003eMinNumber: 起投金额 \u003cbr/\u003eMaxNumber: 最大金额 \u003cbr/\u003eStepNumber: 增幅金额 \u003cbr/\u003eMaxBuySize: 最大购买笔数 \u003cbr/\u003eRefillDaysLimit: 标的物最大期限 \u003cbr/\u003eRefillBidType: 标计息类型 \u003cbr/\u003eRefillBidMark: 标的购买限制 \u003cbr/\u003eRefillRemainNumberMin: 全部符合条件的标 累计小于最小续标金额 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespBids"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/bids/getbidsbyquerycount": {
            "post": {
                "tags": [
                    "/bids"
                ],
                "summary": "获取全部标的物",
                "description": "获取全部标的物",
                "operationId": "/bids/getbidsbyquerycount",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"获取标的物\"\u003cbr/\u003eQuery: 搜索 \u003cbr/\u003eBidType: 标的类型 \u003cbr/\u003eBidMark: 标的分类 \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetBidsByQuery"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "标数据数组数量\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/bids/getbidsbyquerycountnew": {
            "post": {
                "tags": [
                    "/bids"
                ],
                "summary": "查询标的物数量 2017-08-25",
                "description": "查询标的物",
                "operationId": "/bids/getbidsbyquerycountnew",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eQuery: 搜索 \u003cbr/\u003eBidType: 标的类型 \u003cbr/\u003eBidMark: 标的分类 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetBidsByQueryCountNew"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "数量\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/bids/getbidsbyquerynew": {
            "post": {
                "tags": [
                    "/bids"
                ],
                "summary": "查询标的物 2017-08-25",
                "description": "查询标的物",
                "operationId": "/bids/getbidsbyquerynew",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eQuery: 搜索 \u003cbr/\u003eBidType: 标的类型 \u003cbr/\u003eBidMark: 标的分类 \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetBidsByQueryNew"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "标数据数组\u003cbr/\u003eId: 标的id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003eBidType: 标计息类型 标记活期定期 \u003cbr/\u003eBidMark: 标的购买限制 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eRemainNumber: 标剩余金额 \u003cbr/\u003eBidNo: 标订单号 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标详情路径 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eCheckDate: 结算日期 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 \u003cbr/\u003eIsVip: 是计入vip的 \u003cbr/\u003eIsRank: 是计入rank的 \u003cbr/\u003eMark: 标记 \u003cbr/\u003eDays: 实际经过的天数 \u003cbr/\u003eBidRate: 标的物利率 \u003cbr/\u003eIncomeDate: 起息时间 \u003cbr/\u003eBidRateBase: 基础利率 \u003cbr/\u003eBidRateRaise: 加息利率 \u003cbr/\u003eMinNumber: 起投金额 \u003cbr/\u003eMaxNumber: 最大金额 \u003cbr/\u003eStepNumber: 增幅金额 \u003cbr/\u003eMaxBuySize: 最大购买笔数 \u003cbr/\u003eRefillDaysLimit: 标的物最大期限 \u003cbr/\u003eRefillBidType: 标计息类型 \u003cbr/\u003eRefillBidMark: 标的购买限制 \u003cbr/\u003eRefillRemainNumberMin: 全部符合条件的标 累计小于最小续标金额 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespBids"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/bids/getbidsfulld": {
            "post": {
                "tags": [
                    "/bids"
                ],
                "summary": "获取满标的标的物",
                "description": "获取满标的标的物",
                "operationId": "/bids/getbidsfulld",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "标数据数组\u003cbr/\u003eId: 标的id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003eBidType: 标计息类型 标记活期定期 \u003cbr/\u003eBidMark: 标的购买限制 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eRemainNumber: 标剩余金额 \u003cbr/\u003eBidNo: 标订单号 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标详情路径 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eCheckDate: 结算日期 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 \u003cbr/\u003eIsVip: 是计入vip的 \u003cbr/\u003eIsRank: 是计入rank的 \u003cbr/\u003eMark: 标记 \u003cbr/\u003eDays: 实际经过的天数 \u003cbr/\u003eBidRate: 标的物利率 \u003cbr/\u003eIncomeDate: 起息时间 \u003cbr/\u003eBidRateBase: 基础利率 \u003cbr/\u003eBidRateRaise: 加息利率 \u003cbr/\u003eMinNumber: 起投金额 \u003cbr/\u003eMaxNumber: 最大金额 \u003cbr/\u003eStepNumber: 增幅金额 \u003cbr/\u003eMaxBuySize: 最大购买笔数 \u003cbr/\u003eRefillDaysLimit: 标的物最大期限 \u003cbr/\u003eRefillBidType: 标计息类型 \u003cbr/\u003eRefillBidMark: 标的购买限制 \u003cbr/\u003eRefillRemainNumberMin: 全部符合条件的标 累计小于最小续标金额 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespBids"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/bids/getbidsfulldnew": {
            "post": {
                "tags": [
                    "/bids"
                ],
                "summary": "获取所有满标和过期的 2017-08-25",
                "description": "获取所有满标和过期的",
                "operationId": "/bids/getbidsfulldnew",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqOffsetLimit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "标数据数组\u003cbr/\u003eId: 标的id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003eBidType: 标计息类型 标记活期定期 \u003cbr/\u003eBidMark: 标的购买限制 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eRemainNumber: 标剩余金额 \u003cbr/\u003eBidNo: 标订单号 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标详情路径 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eCheckDate: 结算日期 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 \u003cbr/\u003eIsVip: 是计入vip的 \u003cbr/\u003eIsRank: 是计入rank的 \u003cbr/\u003eMark: 标记 \u003cbr/\u003eDays: 实际经过的天数 \u003cbr/\u003eBidRate: 标的物利率 \u003cbr/\u003eIncomeDate: 起息时间 \u003cbr/\u003eBidRateBase: 基础利率 \u003cbr/\u003eBidRateRaise: 加息利率 \u003cbr/\u003eMinNumber: 起投金额 \u003cbr/\u003eMaxNumber: 最大金额 \u003cbr/\u003eStepNumber: 增幅金额 \u003cbr/\u003eMaxBuySize: 最大购买笔数 \u003cbr/\u003eRefillDaysLimit: 标的物最大期限 \u003cbr/\u003eRefillBidType: 标计息类型 \u003cbr/\u003eRefillBidMark: 标的购买限制 \u003cbr/\u003eRefillRemainNumberMin: 全部符合条件的标 累计小于最小续标金额 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespBids"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/bids/getbidsnew": {
            "post": {
                "tags": [
                    "/bids"
                ],
                "summary": "获取所有未满标 2017-08-25",
                "description": "获取所有未满标",
                "operationId": "/bids/getbidsnew",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003e",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetBidsNew"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "标数据数组\u003cbr/\u003eId: 标的id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003eBidType: 标计息类型 标记活期定期 \u003cbr/\u003eBidMark: 标的购买限制 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eRemainNumber: 标剩余金额 \u003cbr/\u003eBidNo: 标订单号 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标详情路径 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eCheckDate: 结算日期 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 \u003cbr/\u003eIsVip: 是计入vip的 \u003cbr/\u003eIsRank: 是计入rank的 \u003cbr/\u003eMark: 标记 \u003cbr/\u003eDays: 实际经过的天数 \u003cbr/\u003eBidRate: 标的物利率 \u003cbr/\u003eIncomeDate: 起息时间 \u003cbr/\u003eBidRateBase: 基础利率 \u003cbr/\u003eBidRateRaise: 加息利率 \u003cbr/\u003eMinNumber: 起投金额 \u003cbr/\u003eMaxNumber: 最大金额 \u003cbr/\u003eStepNumber: 增幅金额 \u003cbr/\u003eMaxBuySize: 最大购买笔数 \u003cbr/\u003eRefillDaysLimit: 标的物最大期限 \u003cbr/\u003eRefillBidType: 标计息类型 \u003cbr/\u003eRefillBidMark: 标的购买限制 \u003cbr/\u003eRefillRemainNumberMin: 全部符合条件的标 累计小于最小续标金额 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespBids"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/bids/getbidssatisfy": {
            "post": {
                "tags": [
                    "/bids"
                ],
                "summary": "获取可以显示未满标的标的物",
                "description": "获取可以显示未满标的标的物",
                "operationId": "/bids/getbidssatisfy",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "标数据数组\u003cbr/\u003eId: 标的id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003eBidType: 标计息类型 标记活期定期 \u003cbr/\u003eBidMark: 标的购买限制 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eRemainNumber: 标剩余金额 \u003cbr/\u003eBidNo: 标订单号 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标详情路径 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eCheckDate: 结算日期 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 \u003cbr/\u003eIsVip: 是计入vip的 \u003cbr/\u003eIsRank: 是计入rank的 \u003cbr/\u003eMark: 标记 \u003cbr/\u003eDays: 实际经过的天数 \u003cbr/\u003eBidRate: 标的物利率 \u003cbr/\u003eIncomeDate: 起息时间 \u003cbr/\u003eBidRateBase: 基础利率 \u003cbr/\u003eBidRateRaise: 加息利率 \u003cbr/\u003eMinNumber: 起投金额 \u003cbr/\u003eMaxNumber: 最大金额 \u003cbr/\u003eStepNumber: 增幅金额 \u003cbr/\u003eMaxBuySize: 最大购买笔数 \u003cbr/\u003eRefillDaysLimit: 标的物最大期限 \u003cbr/\u003eRefillBidType: 标计息类型 \u003cbr/\u003eRefillBidMark: 标的购买限制 \u003cbr/\u003eRefillRemainNumberMin: 全部符合条件的标 累计小于最小续标金额 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespBids"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/bids/getbidssatisfyfulld": {
            "post": {
                "tags": [
                    "/bids"
                ],
                "summary": "获取满标的标的物",
                "description": "获取满标的标的物",
                "operationId": "/bids/getbidssatisfyfulld",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "标数据数组\u003cbr/\u003eId: 标的id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003eBidType: 标计息类型 标记活期定期 \u003cbr/\u003eBidMark: 标的购买限制 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eRemainNumber: 标剩余金额 \u003cbr/\u003eBidNo: 标订单号 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标详情路径 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eCheckDate: 结算日期 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 \u003cbr/\u003eIsVip: 是计入vip的 \u003cbr/\u003eIsRank: 是计入rank的 \u003cbr/\u003eMark: 标记 \u003cbr/\u003eDays: 实际经过的天数 \u003cbr/\u003eBidRate: 标的物利率 \u003cbr/\u003eIncomeDate: 起息时间 \u003cbr/\u003eBidRateBase: 基础利率 \u003cbr/\u003eBidRateRaise: 加息利率 \u003cbr/\u003eMinNumber: 起投金额 \u003cbr/\u003eMaxNumber: 最大金额 \u003cbr/\u003eStepNumber: 增幅金额 \u003cbr/\u003eMaxBuySize: 最大购买笔数 \u003cbr/\u003eRefillDaysLimit: 标的物最大期限 \u003cbr/\u003eRefillBidType: 标计息类型 \u003cbr/\u003eRefillBidMark: 标的购买限制 \u003cbr/\u003eRefillRemainNumberMin: 全部符合条件的标 累计小于最小续标金额 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespBids"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/bids/getfullbids": {
            "post": {
                "tags": [
                    "/bids"
                ],
                "summary": "满标列表",
                "description": "满标列表",
                "operationId": "/bids/getfullbids",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"添加标\"\u003cbr/\u003eBidName: 标名 \u003cbr/\u003eBeginTime: 发布时间 \u003cbr/\u003eEndTime: 发布时间 \u003cbr/\u003eCheckBeginTime: 回款时间 \u003cbr/\u003eCheckEndTime: 汇款时间 \u003cbr/\u003eIncomeBeginTime: 起息时间 \u003cbr/\u003eIncomeEndTime: 起息时间 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqFullBids"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eId: 标的id \u003cbr/\u003eBidType: 标计息类型 标记活期定期 \u003cbr/\u003eBidMark: 标的物分类 \u003cbr/\u003eBidRule: 限制规则 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eRemainNumber: 标剩余金额 \u003cbr/\u003eBidNo: 标订单号 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标详情路径 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eCheckDate: 结算日期 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 修改时间 \u003cbr/\u003eIncomeDate: 起息时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespFullBid"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/bids/updatebid": {
            "post": {
                "tags": [
                    "/bids"
                ],
                "summary": "修改标的物",
                "description": "修改标的物",
                "operationId": "/bids/updatebid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"添加标\"\u003cbr/\u003eBidId: 标id \u003cbr/\u003eBidMark: 标的购买限制 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标的详情地址 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eCheckDate: 结算日期 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eMinNumber: 起投金额 \u003cbr/\u003eMaxNumber: 最大金额 \u003cbr/\u003eStepNumber: 增幅金额 \u003cbr/\u003eMaxBuySize: 最大购买笔数 \u003cbr/\u003eRefillDaysLimit: 标的物最大期限 \u003cbr/\u003eRefillBidType: 标计息类型 \u003cbr/\u003eRefillBidMark: 标的购买限制 \u003cbr/\u003eRefillRemainNumberMin: 全部符合条件的标 累计小于最小续标金额 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUpdateBid"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "标的id\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/bids/updatebidnew": {
            "post": {
                "tags": [
                    "/bids"
                ],
                "summary": "修改标 2017-08-25",
                "description": "修改标",
                "operationId": "/bids/updatebidnew",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eBidId: 标id \u003cbr/\u003eBidMark: 标的购买限制 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标的详情地址 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eMinNumber: 起投金额 \u003cbr/\u003eMaxNumber: 最大金额 \u003cbr/\u003eStepNumber: 增幅金额 \u003cbr/\u003eMaxBuySize: 最大购买笔数 \u003cbr/\u003eRefillDaysLimit: 标的物最大期限 \u003cbr/\u003eRefillBidType: 标计息类型 \u003cbr/\u003eRefillBidMark: 标的购买限制 \u003cbr/\u003eRefillRemainNumberMin: 全部符合条件的标 累计小于最小续标金额 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUpdateBidNew"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "标的id\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/bids/updatebidpool": {
            "post": {
                "tags": [
                    "/bids"
                ],
                "summary": "添加标的物池",
                "description": "添加标的物池",
                "operationId": "/bids/updatebidpool",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"添加活期标\"\u003cbr/\u003eBidId: 标id \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eBidInfoPath: 标的详情地址 \u003cbr/\u003eCheckDate: 结算日期 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUpdateBidPool"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "标的id\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/bids/updatebidsatisfy": {
            "post": {
                "tags": [
                    "/bids"
                ],
                "summary": "修改标的物 满标计息",
                "description": "修改标的物 满标计息",
                "operationId": "/bids/updatebidsatisfy",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"添加标\"\u003cbr/\u003eBidId: 标id \u003cbr/\u003eBidMark: 标的购买限制 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标的详情地址 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eCheckDate: 结算日期 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eMinNumber: 起投金额 \u003cbr/\u003eMaxNumber: 最大金额 \u003cbr/\u003eStepNumber: 增幅金额 \u003cbr/\u003eMaxBuySize: 最大购买笔数 \u003cbr/\u003eRefillDaysLimit: 标的物最大期限 \u003cbr/\u003eRefillBidType: 标计息类型 \u003cbr/\u003eRefillBidMark: 标的购买限制 \u003cbr/\u003eRefillRemainNumberMin: 全部符合条件的标 累计小于最小续标金额 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUpdateBidSatisfy"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "标的id\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/broadcast/list": {
            "post": {
                "tags": [
                    "/broadcast"
                ],
                "summary": "公告消息列表",
                "description": "公告消息列表",
                "operationId": "/broadcast/list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "消息列表\u003cbr/\u003eId: 广播id \u003cbr/\u003eState: 状态 \u003cbr/\u003eMessage: 信息 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.Broadcast"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/chargerecord/getchargerecord": {
            "post": {
                "tags": [
                    "/chargerecord"
                ],
                "summary": "获取用户的充值记录",
                "description": "获取用户的充值记录",
                "operationId": "/chargerecord/getchargerecord",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户充值记录数据\u003cbr/\u003eId: 充值记录id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eNumber: 支付金额 \u003cbr/\u003ePayWay: 支付方式 \u003cbr/\u003eBankcardId: 银行卡id \u003cbr/\u003ePayNo: 订单号 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.ChargeRecord"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/chargerecord/getchargerecordbyno": {
            "post": {
                "tags": [
                    "/chargerecord"
                ],
                "summary": "获取用户的充值记录 根据支付订单",
                "description": "获取用户的充值记录 根据支付订单",
                "operationId": "/chargerecord/getchargerecordbyno",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"支付订单\"\u003cbr/\u003ePayNo: 支付订单号 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetChargeRecordByPayNo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "充值记录数据\u003cbr/\u003eId: 充值记录id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eNumber: 支付金额 \u003cbr/\u003ePayWay: 支付方式 \u003cbr/\u003eBankcardId: 银行卡id \u003cbr/\u003ePayNo: 订单号 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.ChargeRecord"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/chargerecord/getchargerecordbyuid": {
            "post": {
                "tags": [
                    "/chargerecord"
                ],
                "summary": "获取用户的充值记录",
                "description": "获取用户的充值记录",
                "operationId": "/chargerecord/getchargerecordbyuid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"支付订单\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserIdOffsetLimit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "充值记录数据\u003cbr/\u003eId: 充值记录id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eNumber: 支付金额 \u003cbr/\u003ePayWay: 支付方式 \u003cbr/\u003eBankcardId: 银行卡id \u003cbr/\u003ePayNo: 订单号 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.ChargeRecord"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/chargerecord/getchargerecordstatistics": {
            "post": {
                "tags": [
                    "/chargerecord"
                ],
                "summary": "获取用户的充值记录",
                "description": "获取用户的充值记录",
                "operationId": "/chargerecord/getchargerecordstatistics",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"支付订单\"\u003cbr/\u003eUserName: 实名 \u003cbr/\u003eAccount: 账号 \u003cbr/\u003eBeginTime: 开始时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eOffset: 偏移 \u003cbr/\u003eLimit: 限制 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserTime"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "充值记录数据\u003cbr/\u003eAccount: 账号 \u003cbr/\u003eUserName: 实名 \u003cbr/\u003eNumber: 充值金额 \u003cbr/\u003ePayWay: 支付渠道 \u003cbr/\u003ePayNo: 订单号 \u003cbr/\u003eBankCode: 银行编码 \u003cbr/\u003eBankName: 银行名 \u003cbr/\u003eCard: 卡号 \u003cbr/\u003eCreateTime: 充值时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespChargeRecordStatistics"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/chargerecord/getuserchargerecord": {
            "post": {
                "tags": [
                    "/chargerecord"
                ],
                "summary": "获取用户的充值记录数量",
                "description": "获取用户的充值记录数量",
                "operationId": "/chargerecord/getuserchargerecord",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户名\"\u003cbr/\u003eQuery: 查询 \u003cbr/\u003eTimeBegin: 开始时间 \u003cbr/\u003eTimeEnd: 结束时间 \u003cbr/\u003eNumberBegin: 最小金额 \u003cbr/\u003eNumberEnd: 最大金额 \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserAndPage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户充值记录数据\u003cbr/\u003eAccount: 电话号码 \u003cbr/\u003eUserName: 真实姓名 \u003cbr/\u003eId: \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eNumber: 支付金额 \u003cbr/\u003ePayWay: 支付方式 \u003cbr/\u003eBankcardId: 银行卡id \u003cbr/\u003eCard: 银行卡号 \u003cbr/\u003ePayNo: 订单号 \u003cbr/\u003eTpOrderNo: 第三方订单号 \u003cbr/\u003eCreateTime: ",
                        "schema": {
                            "$ref": "#/definitions/models.ResqUserChargeRecord"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/config/configs": {
            "post": {
                "tags": [
                    "/config"
                ],
                "summary": "获取配置文件",
                "description": "获取配置文件",
                "operationId": "/config/configs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "配置文件\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/config/getconfig": {
            "post": {
                "tags": [
                    "/config"
                ],
                "summary": "获取配置文件",
                "description": "获取配置文件",
                "operationId": "/config/getconfig",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "配置文件\u003cbr/\u003eId: 配置id \u003cbr/\u003eKey: 键 \u003cbr/\u003eValue: 值 \u003cbr/\u003eDesc: 描述 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.Config"
                        }
                    },
                    "202": {
                        "description": "各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/coupon/getcoupon": {
            "post": {
                "tags": [
                    "/coupon"
                ],
                "summary": "返回优惠券类型",
                "description": "返回优惠券类型",
                "operationId": "/coupon/getcoupon",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"优惠券类型id\"\u003cbr/\u003eCouponId: 优惠券类型 id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqCouponId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回优惠券类型\u003cbr/\u003eId: 优惠券类型id \u003cbr/\u003eCouponType: 优惠券类型  0 | CouponTypeDiscount | CouponTypeRaise \u003cbr/\u003eCouponRule: 优惠券规则 \u003cbr/\u003eCouponName: 优惠券描述 \u003cbr/\u003eCouponDesc: 优惠券信息 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.Coupon"
                        }
                    },
                    "202": {
                        "description": "操作失败各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/dividendsub/adddividendsub": {
            "post": {
                "tags": [
                    "/dividendsub"
                ],
                "summary": "添加标的利息区间",
                "description": "添加标的利息区间",
                "operationId": "/dividendsub/adddividendsub",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"利息区间id\"\u003cbr/\u003eBidId: 标的id \u003cbr/\u003eRate: 利率 \u003cbr/\u003eBeginDate: 开始日期 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqAddDividendSub"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "利息区间id\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/dividendsub/getdividendsub": {
            "post": {
                "tags": [
                    "/dividendsub"
                ],
                "summary": "获取标的利息区间",
                "description": "获取标的利息区间",
                "operationId": "/dividendsub/getdividendsub",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"利息区间\"\u003cbr/\u003eBidId: 标的id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetDividendSub"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "利息区间id\u003cbr/\u003e",
                        "schema": {
                            "$ref": "#/definitions/models.DividendSubs"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/feedback/record": {
            "post": {
                "tags": [
                    "/feedback"
                ],
                "summary": "反馈记录",
                "description": "反馈记录",
                "operationId": "/feedback/record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBeginTime: 开始时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eOffset: 偏移 \u003cbr/\u003eLimit: 限制 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserTimeRangeOffsetLimit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eId: \u003cbr/\u003eAccount: \u003cbr/\u003eUserName: \u003cbr/\u003eContent: \u003cbr/\u003eCreateTime: ",
                        "schema": {
                            "$ref": "#/definitions/models.RespFeedbackRecord"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/feedback/submit": {
            "post": {
                "tags": [
                    "/feedback"
                ],
                "summary": "反馈提交",
                "description": "反馈提交",
                "operationId": "/feedback/submit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eContent: 内容 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqFeedback"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/finance/getreturnrecord": {
            "post": {
                "tags": [
                    "/finance"
                ],
                "summary": "获取某天回款数据",
                "description": "获取某天回款数据",
                "operationId": "/finance/getreturnrecord",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"日期\"\u003cbr/\u003eBeginTime: 起始日期 \u003cbr/\u003eEndTime: 结束日期 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBeginEndDate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eUserName: \u003cbr/\u003eBidName: \u003cbr/\u003eDrawNumber: \u003cbr/\u003eEarningsNumber: \u003cbr/\u003eCreateTime: ",
                        "schema": {
                            "$ref": "#/definitions/models.ResqReturnRecordInfo"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/message/add": {
            "post": {
                "tags": [
                    "/message"
                ],
                "summary": "添加消息",
                "description": "添加消息",
                "operationId": "/message/add",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"参数\"\u003cbr/\u003eTitle: 标题 \u003cbr/\u003eMsg: 消息 \u003cbr/\u003eUrl: 跳转链接 \u003cbr/\u003eBeginTime: 开始时间 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqMessageAdd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/message/addallusermessage": {
            "post": {
                "tags": [
                    "/message"
                ],
                "summary": "给所有用户发送消息",
                "description": "给所有用户发送消息",
                "operationId": "/message/addallusermessage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"参数\"\u003cbr/\u003eTitle: 标题 \u003cbr/\u003eMsg: 消息 \u003cbr/\u003eUrl: 跳转链接 \u003cbr/\u003eBeginTime: 发布时间 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqAllMessageAdd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "消息列表\u003cbr/\u003e",
                        "schema": {
                            "$ref": "#/definitions/apim.Message"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/message/addmessagebyaccount": {
            "post": {
                "tags": [
                    "/message"
                ],
                "summary": "获取消息详情",
                "description": "获取消息详情",
                "operationId": "/message/addmessagebyaccount",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"参数\"\u003cbr/\u003eTitle: 标题 \u003cbr/\u003eMsg: 消息 \u003cbr/\u003eUrl: 跳转链接 \u003cbr/\u003eBeginTime: 发布时间 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqMessageAddByAccount"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "消息列表\u003cbr/\u003e",
                        "schema": {
                            "$ref": "#/definitions/int"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/message/getmessageinfo": {
            "post": {
                "tags": [
                    "/message"
                ],
                "summary": "获取消息详情",
                "description": "获取消息详情",
                "operationId": "/message/getmessageinfo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"参数\"\u003cbr/\u003eMessageId: 消息id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqMessageInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "消息列表\u003cbr/\u003e",
                        "schema": {
                            "$ref": "#/definitions/apim.Message"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/message/getmessages": {
            "post": {
                "tags": [
                    "/message"
                ],
                "summary": "获取消息",
                "description": "获取消息",
                "operationId": "/message/getmessages",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"参数\"\u003cbr/\u003eMessageId: id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eState: 状态 \u003cbr/\u003eMsgType: 消息类型 \u003cbr/\u003eTitle: 标题 \u003cbr/\u003eBeginTime: 开始时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetMessages"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "消息列表\u003cbr/\u003eMessageId: id \u003cbr/\u003eTitle: 标题 \u003cbr/\u003eMsg: 内容 \u003cbr/\u003eMsgType: 消息类型 \u003cbr/\u003eReadCount: 已读数量 \u003cbr/\u003eTotalCount: 总数 \u003cbr/\u003eBeginTime: 显示时间 \u003cbr/\u003eCreateTime: 发布时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespMessage"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/message/list": {
            "post": {
                "tags": [
                    "/message"
                ],
                "summary": "消息列表",
                "description": "delete the Message",
                "operationId": "/message/list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"参数\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqMessageList"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "消息列表\u003cbr/\u003eId: 信息id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eState: 信息状态 \u003cbr/\u003eTitle: 消息标题 \u003cbr/\u003eMsg: 信息 \u003cbr/\u003eUrl: 点击后跳转链接 \u003cbr/\u003eBeginTime: \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/message/read": {
            "post": {
                "tags": [
                    "/message"
                ],
                "summary": "设置消息已读",
                "description": "设置消息已读",
                "operationId": "/message/read",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"参数\"\u003cbr/\u003eId: 消息id \u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqMessageRead"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/message/remain": {
            "post": {
                "tags": [
                    "/message"
                ],
                "summary": "未读消息条数",
                "description": "未读消息条数",
                "operationId": "/message/remain",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"参数\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqMessageCount"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "未读消息条数\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/message/updatemessage": {
            "post": {
                "tags": [
                    "/message"
                ],
                "summary": "修改消息详情",
                "description": "修改消息详情",
                "operationId": "/message/updatemessage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"参数\"\u003cbr/\u003eMessageId: 消息id \u003cbr/\u003eTitle: 标题 \u003cbr/\u003eMsg: 消息 \u003cbr/\u003eUrl: 跳转链接 \u003cbr/\u003eBeginTime: 发布时间 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqMessageUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "消息列表\u003cbr/\u003eId: 信息id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eState: 信息状态 \u003cbr/\u003eTitle: 消息标题 \u003cbr/\u003eMsg: 信息 \u003cbr/\u003eUrl: 点击后跳转链接 \u003cbr/\u003eBeginTime: \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/platform/addplatform": {
            "post": {
                "tags": [
                    "/platform"
                ],
                "summary": "渠道状态修改",
                "description": "渠道状态修改",
                "operationId": "/platform/addplatform",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eState: 状态 \u003cbr/\u003ePlatformId: 渠道id \u003cbr/\u003ePackCode: 包名 \u003cbr/\u003eName: 渠道名称 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqAddPlatform"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/platform/getplatformbycode": {
            "post": {
                "tags": [
                    "/platform"
                ],
                "summary": "获取包名对应渠道",
                "description": "获取包名对应渠道",
                "operationId": "/platform/getplatformbycode",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003ePackCode: 包名 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqPlatformPack"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eId: \u003cbr/\u003ePlatformId: \u003cbr/\u003eState: \u003cbr/\u003ePackCode: \u003cbr/\u003eName: \u003cbr/\u003eCreateTime: ",
                        "schema": {
                            "$ref": "#/definitions/models.Platform"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/platform/getplatformbyquery": {
            "post": {
                "tags": [
                    "/platform"
                ],
                "summary": "查询渠道",
                "description": "查询渠道",
                "operationId": "/platform/getplatformbyquery",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eState: 状态 \u003cbr/\u003ePlatformId: 渠道id \u003cbr/\u003ePackCode: 包名 \u003cbr/\u003eName: 渠道名称 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqPlatform"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eId: \u003cbr/\u003ePlatformId: \u003cbr/\u003eState: \u003cbr/\u003ePackCode: \u003cbr/\u003eName: \u003cbr/\u003eCreateTime: ",
                        "schema": {
                            "$ref": "#/definitions/models.Platform"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/platform/getplatformfoundstatistics": {
            "post": {
                "tags": [
                    "/platform"
                ],
                "summary": "渠道用户投资统计",
                "description": "渠道用户投资统计",
                "operationId": "/platform/getplatformfoundstatistics",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003ePlatformId: 渠道id \u003cbr/\u003eBeginTime: 时间 \u003cbr/\u003eEndTime: 时间 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqPlatformStatistics"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/platform/getplatformpurchasestatistics": {
            "post": {
                "tags": [
                    "/platform"
                ],
                "summary": "渠道用户投资统计",
                "description": "渠道用户投资统计",
                "operationId": "/platform/getplatformpurchasestatistics",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003ePlatformId: 渠道id \u003cbr/\u003eBeginTime: 时间 \u003cbr/\u003eEndTime: 时间 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqPlatformStatistics"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/platform/getplatformuserinfostatistics": {
            "post": {
                "tags": [
                    "/platform"
                ],
                "summary": "渠道用户统计",
                "description": "渠道用户统计",
                "operationId": "/platform/getplatformuserinfostatistics",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003ePlatformId: 渠道id \u003cbr/\u003eBeginTime: 时间 \u003cbr/\u003eEndTime: 时间 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqPlatformStatistics"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/platform/updateplatform": {
            "post": {
                "tags": [
                    "/platform"
                ],
                "summary": "渠道状态修改",
                "description": "渠道状态修改",
                "operationId": "/platform/updateplatform",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eId: id \u003cbr/\u003eState: 状态 \u003cbr/\u003ePlatformId: 渠道id \u003cbr/\u003ePackCode: 包名 \u003cbr/\u003eName: 渠道名称 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUpdatePlatform"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/platform/updatestate": {
            "post": {
                "tags": [
                    "/platform"
                ],
                "summary": "渠道状态修改",
                "description": "渠道状态修改",
                "operationId": "/platform/updatestate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eId: 渠道设置id \u003cbr/\u003eState: 状态 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqPlatformState"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/purchaserecord/countpurchaserecordnew": {
            "post": {
                "tags": [
                    "/purchaserecord"
                ],
                "summary": "投资记录 2017-09-17",
                "description": "投资记录",
                "operationId": "/purchaserecord/countpurchaserecordnew",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserName: 姓名 \u003cbr/\u003eAccount: 账号 \u003cbr/\u003eUserId: uid \u003cbr/\u003eState: 投资状态 \u003cbr/\u003eOffset: 起始偏移 \u003cbr/\u003eLimit: 限制取多少条 \u003cbr/\u003eBeginTime: 投资起始日期 \u003cbr/\u003eEndTime: 投资结束日期 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqPurchaseRecordNew"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e",
                        "schema": {
                            "$ref": "#/definitions/int"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/purchaserecord/getcurrentpurchaserecord": {
            "post": {
                "tags": [
                    "/purchaserecord"
                ],
                "summary": "活期购买查询",
                "description": "活期购买查询",
                "operationId": "/purchaserecord/getcurrentpurchaserecord",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBeginTime: 开始时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eType: \"CURRENT\"活期 \"FIXED\"定期 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqPurchaseRecord"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eId: 购买池id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eNumber: 购买金额 \u003cbr/\u003ePayNumber: 支付金额 \u003cbr/\u003eRemainNumber: 持有金额 \u003cbr/\u003eStatus: 持有状态 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.PurchaseRecordPool"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/purchaserecord/getpurchasefullcount": {
            "post": {
                "tags": [
                    "/purchaserecord"
                ],
                "summary": "获取满足条件购标统计",
                "description": "获取满足条件购标统计",
                "operationId": "/purchaserecord/getpurchasefullcount",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eNumber: 提现金额 \u003cbr/\u003eBeginTime: 开始时间 \u003cbr/\u003eEndTime: 结束时间 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetPurchaseFullCount"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eFullCount: \u003cbr/\u003eEveryFullCount: \u003cbr/\u003ePurchaseNumber: ",
                        "schema": {
                            "$ref": "#/definitions/models.RespPurchaseFullCount"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/purchaserecord/getpurchaseinfo": {
            "post": {
                "tags": [
                    "/purchaserecord"
                ],
                "summary": "购买详情",
                "description": "购买详情",
                "operationId": "/purchaserecord/getpurchaseinfo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eBeginTime: 起始日期 \u003cbr/\u003eEndTime: 结束日期 \u003cbr/\u003eUserId: 用户id \u003cbr/\u003ePlatform: 渠道id \u003cbr/\u003eAccount: 账号 \u003cbr/\u003eUserName: 用户真实姓名 \u003cbr/\u003eBidName: 标名 \u003cbr/\u003eBidNo: 标的编号 \u003cbr/\u003eOffset: 起始偏移 \u003cbr/\u003eLimit: 限制取多少条 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqPurchaseStatisticsRecord"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eAccount: 账号 \u003cbr/\u003eInviter: 邀请人id \u003cbr/\u003eInviterAccount: 邀请人账号 \u003cbr/\u003eBidId: 标id \u003cbr/\u003eUserName: 姓名 \u003cbr/\u003ePlatform: 渠道 \u003cbr/\u003ePlatformName: 渠道名称 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidNo: 标编号 \u003cbr/\u003eDaysLimit: 标期限 \u003cbr/\u003eBidRateShow: 标利率 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eRemainNumber: 标剩余金额 \u003cbr/\u003eCheckDate: 回款日期 \u003cbr/\u003eEndTime: 止卖时间 \u003cbr/\u003eNumber: 投资金额 \u003cbr/\u003ePayNumber: 实付金额 \u003cbr/\u003eCouponRules: 券规则 \u003cbr/\u003eCouponName: 券名称 \u003cbr/\u003eCreateTime: 投资时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespStatisticsPurchaseInfo"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/purchaserecord/getpurchaseinfototal": {
            "post": {
                "tags": [
                    "/purchaserecord"
                ],
                "summary": "购买详情统计",
                "description": "购买详情统计",
                "operationId": "/purchaserecord/getpurchaseinfototal",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eBeginTime: 起始日期 \u003cbr/\u003eEndTime: 结束日期 \u003cbr/\u003eDaysLimitMax: 天数最大 \u003cbr/\u003eDaysLimitMin: 天数最小 \u003cbr/\u003eUserId: 用户id \u003cbr/\u003ePlatform: 渠道id \u003cbr/\u003eAccount: 账号 \u003cbr/\u003eUserName: 用户真实姓名 \u003cbr/\u003eBidName: 标名 \u003cbr/\u003eBidNo: 标的编号 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetPurchaseInfoTotal"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/purchaserecord/getpurchaserecordbyuserid": {
            "post": {
                "tags": [
                    "/purchaserecord"
                ],
                "summary": "用户购买记录 根据购买时间排序",
                "description": "用户购买记录 根据购买时间排序",
                "operationId": "/purchaserecord/getpurchaserecordbyuserid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eOffset: 起始偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetPurchaseRecordByUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eId: \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003ePurchaseRecordPoolId: 购买记录池id \u003cbr/\u003eNumber: 购买金额 \u003cbr/\u003ePayNumber: 支付金额 \u003cbr/\u003eRemainNumber: 剩余金额 \u003cbr/\u003eCouponRules: 购买时使用的券 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.PurchaseRecord"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/purchaserecord/getpurchaserecordnew": {
            "post": {
                "tags": [
                    "/purchaserecord"
                ],
                "summary": "投资记录 2017-09-17",
                "description": "投资记录",
                "operationId": "/purchaserecord/getpurchaserecordnew",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserName: 姓名 \u003cbr/\u003eAccount: 账号 \u003cbr/\u003eUserId: uid \u003cbr/\u003eState: 投资状态 \u003cbr/\u003eOffset: 起始偏移 \u003cbr/\u003eLimit: 限制取多少条 \u003cbr/\u003eBeginTime: 投资起始日期 \u003cbr/\u003eEndTime: 投资结束日期 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqPurchaseRecordNew"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003ePurchaseRecordId: 投资记录id \u003cbr/\u003eBidId: 标id \u003cbr/\u003ePurchaseRecordPoolId: 活期池id \u003cbr/\u003eNumber: 投资金额 \u003cbr/\u003ePayNumber: 实付金额 \u003cbr/\u003ePurchaseRemainNumber: 活期剩余金额 \u003cbr/\u003eCouponRules: 使用的券 \u003cbr/\u003eCreateTime: 购买时间 \u003cbr/\u003eUserName: 实名 \u003cbr/\u003eAccount: 账号 \u003cbr/\u003eBidName: 标名 \u003cbr/\u003eDaysLimit: 期限天数 \u003cbr/\u003eBidRateShow: 标利率展示 \u003cbr/\u003eRemainNumber: 标剩余金额 \u003cbr/\u003eBeginTime: 标起售时间 \u003cbr/\u003eEndTime: 标结束时间 \u003cbr/\u003eCheckDate: 回款时间 \u003cbr/\u003eRate: 标的利率 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespPurchaseRecordNew"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/purchaserecord/getpurchaserecordssum": {
            "post": {
                "tags": [
                    "/purchaserecord"
                ],
                "summary": "统计放款和回款数据",
                "description": "统计放款和回款数据",
                "operationId": "/purchaserecord/getpurchaserecordssum",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eLoansNumber: 当前还没回款的金额 \u003cbr/\u003eLoansCount: 当前还没回款的笔数 \u003cbr/\u003eReceiptNumber: 当前已经回款的金额 \u003cbr/\u003eReceiptCount: 当前已经回款的笔数 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespPurchaseRecordsSum"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/purchaserecord/getpurchasestatement": {
            "post": {
                "tags": [
                    "/purchaserecord"
                ],
                "summary": "投资汇总",
                "description": "投资汇总",
                "operationId": "/purchaserecord/getpurchasestatement",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eBeginTime: 起始日期 \u003cbr/\u003eEndTime: 结束日期 \u003cbr/\u003eOffset: 起始偏移 \u003cbr/\u003eLimit: 限制取多少条 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqDateOffsetLimit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/purchaserecord/gettimetotal": {
            "post": {
                "tags": [
                    "/purchaserecord"
                ],
                "summary": "投资汇总",
                "description": "投资汇总,一段时间内的销量",
                "operationId": "/purchaserecord/gettimetotal",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eBeginTime: 起始日期 \u003cbr/\u003eEndTime: 结束日期 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBeginEndDate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eNumber: 投资金额 \u003cbr/\u003ePayNumber: 实付 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespPurchaseTotal"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/purchaserecord/getunreturntotal": {
            "post": {
                "tags": [
                    "/purchaserecord"
                ],
                "summary": "未回款统计",
                "description": "未回款统计",
                "operationId": "/purchaserecord/getunreturntotal",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eBeginTime: 起始日期 \u003cbr/\u003eEndTime: 结束日期 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBeginEndDate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/purchaserecord/getuserpurchasestatement": {
            "post": {
                "tags": [
                    "/purchaserecord"
                ],
                "summary": "获取用户的投资情况",
                "description": "获取用户的投资情况",
                "operationId": "/purchaserecord/getuserpurchasestatement",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"购买记录\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eNewNumber: 新手标投资金额 \u003cbr/\u003eNewCount: 新手标投资次数 \u003cbr/\u003eTotalNumber: 总投资金额 \u003cbr/\u003eTotalCount: 总投资次数 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespUserPurchaseStatement"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/purchaserecord/platformpurchaserecord": {
            "post": {
                "tags": [
                    "/purchaserecord"
                ],
                "summary": "投资记录导出",
                "description": "投资记录导出",
                "operationId": "/purchaserecord/platformpurchaserecord",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eBeginTime: 投资起始日期 \u003cbr/\u003eEndTime: 投资结束日期 \u003cbr/\u003eUserName: 姓名 \u003cbr/\u003eAccount: 账号 \u003cbr/\u003eOffset: 起始偏移 \u003cbr/\u003eLimit: 限制取多少条 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqPlatformPurchase"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eAccount: 账号 \u003cbr/\u003eUserName: 实名 \u003cbr/\u003ePlatform: 渠道 \u003cbr/\u003eInviter: 邀请人id \u003cbr/\u003eInviterName: 邀请人实名 \u003cbr/\u003eInviterAccount: 邀请人账号 \u003cbr/\u003eRegisterTime: 注册时间 \u003cbr/\u003eIdCard: 身份证号 \u003cbr/\u003eBankcard: 用户银行卡号 \u003cbr/\u003ePurchaseCount: 购买次数 \u003cbr/\u003eBidName: 标名 \u003cbr/\u003eDaysLimit: 期限天数 \u003cbr/\u003eNumber: 投资金额 \u003cbr/\u003ePayNumber: 实付金额 \u003cbr/\u003eCreateTime: 购买时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespPurchaseExport"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/purchaserecord/purchaseexport": {
            "post": {
                "tags": [
                    "/purchaserecord"
                ],
                "summary": "投资记录导出",
                "description": "投资记录导出",
                "operationId": "/purchaserecord/purchaseexport",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eRegisterBeginTime: 注册起始时间 \u003cbr/\u003eRegisterEndTime: 注册截至时间 \u003cbr/\u003eBeginTime: 投资起始日期 \u003cbr/\u003eEndTime: 投资结束日期 \u003cbr/\u003ePlatform: 渠道id \u003cbr/\u003eOffset: 起始偏移 \u003cbr/\u003eLimit: 限制取多少条 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqPurchaseExport"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eAccount: 账号 \u003cbr/\u003eUserName: 实名 \u003cbr/\u003ePlatform: 渠道 \u003cbr/\u003eInviter: 邀请人id \u003cbr/\u003eInviterName: 邀请人实名 \u003cbr/\u003eInviterAccount: 邀请人账号 \u003cbr/\u003eRegisterTime: 注册时间 \u003cbr/\u003eIdCard: 身份证号 \u003cbr/\u003eBankcard: 用户银行卡号 \u003cbr/\u003ePurchaseCount: 购买次数 \u003cbr/\u003eBidName: 标名 \u003cbr/\u003eDaysLimit: 期限天数 \u003cbr/\u003eNumber: 投资金额 \u003cbr/\u003ePayNumber: 实付金额 \u003cbr/\u003eCreateTime: 购买时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespPurchaseExport"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/rank/getranking": {
            "post": {
                "tags": [
                    "/rank"
                ],
                "summary": "获取一段时间内的累投排行榜",
                "description": "获取一段时间内的累投排行榜",
                "operationId": "/rank/getranking",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\" \"\u003cbr/\u003eBeginTime: 开始时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetRanking"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eRank: 分数 \u003cbr/\u003eNickName: 昵称 \u003cbr/\u003eUserName: 真实姓名 \u003cbr/\u003eAccount: 账号 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespGetRanking"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/rank/getrankingbyday": {
            "post": {
                "tags": [
                    "/rank"
                ],
                "summary": "获取日累投的排行榜",
                "description": "获取日累投的排行榜",
                "operationId": "/rank/getrankingbyday",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\" \"\u003cbr/\u003eQuery: 搜索 \u003cbr/\u003eBeginTime: 开始时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetRankingByDay"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eRank: 分数 \u003cbr/\u003eNickName: 昵称 \u003cbr/\u003eUserName: 真实姓名 \u003cbr/\u003eAccount: 账号 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespGetRanking"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/record/balance": {
            "post": {
                "tags": [
                    "/record"
                ],
                "summary": "余额变更记录",
                "description": "余额变更记录",
                "operationId": "/record/balance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserIdOffsetLimit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "余额变更记录\u003cbr/\u003eNumber: 变动的金额 \u003cbr/\u003eCreateTime: 变动的时间 \u003cbr/\u003eMsg: 变动描述 ",
                        "schema": {
                            "$ref": "#/definitions/models.ResqBalanceChange"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/record/getpurchase": {
            "post": {
                "tags": [
                    "/record"
                ],
                "summary": "获取用户购买记录",
                "description": "获取用户购买记录",
                "operationId": "/record/getpurchase",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"购买记录id\"\u003cbr/\u003ePurchaseRecordId: 购买记录id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqPurchaseRecordId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "购买池记录数组\u003cbr/\u003eId: 标的id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003eBidType: 标计息类型 标记活期定期 \u003cbr/\u003eBidMark: 标的购买限制 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eRemainNumber: 标剩余金额 \u003cbr/\u003eBidNo: 标订单号 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标详情路径 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eCheckDate: 结算日期 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 \u003cbr/\u003eIsVip: 是计入vip的 \u003cbr/\u003eIsRank: 是计入rank的 \u003cbr/\u003eMark: 标记 \u003cbr/\u003eDays: 实际经过的天数 \u003cbr/\u003eBidRate: 标的物利率 \u003cbr/\u003eIncomeDate: 起息时间 \u003cbr/\u003eBidRateBase: 基础利率 \u003cbr/\u003eBidRateRaise: 加息利率 \u003cbr/\u003eNumber: 金额 \u003cbr/\u003eDrawNumber: 持有金额 \u003cbr/\u003eEarningsNumber: 购买金额的收益 \u003cbr/\u003eYesterEarningsNumber: 昨日收益 \u003cbr/\u003eExpireEarningsNumber: 预计到期收益 \u003cbr/\u003eExpireDays: 预计天数 废弃 \u003cbr/\u003ePayNumber: 支付金额 \u003cbr/\u003ePurchaseRecordId: 购买记录id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eIsIncome: 是否起息 \u003cbr/\u003eCouponMsg: 优惠券使用描述 \u003cbr/\u003eMinNumber: 起投金额 \u003cbr/\u003eMaxNumber: 最大金额 \u003cbr/\u003eStepNumber: 增幅金额 \u003cbr/\u003eMaxBuySize: 最大购买笔数 \u003cbr/\u003eRefillDaysLimit: 标的物最大期限 \u003cbr/\u003eRefillBidType: 标计息类型 \u003cbr/\u003eRefillBidMark: 标的购买限制 \u003cbr/\u003eRefillRemainNumberMin: 全部符合条件的标 累计小于最小续标金额 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespPurchaseRecord"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/record/getpurchaseexistbybidid": {
            "post": {
                "tags": [
                    "/record"
                ],
                "summary": "获取标的还有余额的购买记录",
                "description": "获取标的还有余额的购买记录",
                "operationId": "/record/getpurchaseexistbybidid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"标的id\"\u003cbr/\u003eBidId: 标的id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBidId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "操作成功\u003cbr/\u003eId: 标的id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003eBidType: 标计息类型 标记活期定期 \u003cbr/\u003eBidMark: 标的购买限制 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eRemainNumber: 标剩余金额 \u003cbr/\u003eBidNo: 标订单号 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标详情路径 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eCheckDate: 结算日期 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 \u003cbr/\u003eIsVip: 是计入vip的 \u003cbr/\u003eIsRank: 是计入rank的 \u003cbr/\u003eMark: 标记 \u003cbr/\u003eDays: 实际经过的天数 \u003cbr/\u003eBidRate: 标的物利率 \u003cbr/\u003eIncomeDate: 起息时间 \u003cbr/\u003eBidRateBase: 基础利率 \u003cbr/\u003eBidRateRaise: 加息利率 \u003cbr/\u003eNumber: 金额 \u003cbr/\u003eDrawNumber: 持有金额 \u003cbr/\u003eEarningsNumber: 购买金额的收益 \u003cbr/\u003eYesterEarningsNumber: 昨日收益 \u003cbr/\u003eExpireEarningsNumber: 预计到期收益 \u003cbr/\u003eExpireDays: 预计天数 废弃 \u003cbr/\u003ePayNumber: 支付金额 \u003cbr/\u003ePurchaseRecordId: 购买记录id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eIsIncome: 是否起息 \u003cbr/\u003eCouponMsg: 优惠券使用描述 \u003cbr/\u003eMinNumber: 起投金额 \u003cbr/\u003eMaxNumber: 最大金额 \u003cbr/\u003eStepNumber: 增幅金额 \u003cbr/\u003eMaxBuySize: 最大购买笔数 \u003cbr/\u003eRefillDaysLimit: 标的物最大期限 \u003cbr/\u003eRefillBidType: 标计息类型 \u003cbr/\u003eRefillBidMark: 标的购买限制 \u003cbr/\u003eRefillRemainNumberMin: 全部符合条件的标 累计小于最小续标金额 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespPurchaseRecord"
                        }
                    },
                    "202": {
                        "description": "操作失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/record/getpurchaseexistbypoolid": {
            "post": {
                "tags": [
                    "/record"
                ],
                "summary": "获取用户对应购买池的购买记录包括余额为0的记录",
                "description": "获取用户对应购买池的购买记录包括余额为0的记录",
                "operationId": "/record/getpurchaseexistbypoolid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"购买池id\"\u003cbr/\u003ePurchaseRecordPoolId: ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqPurchaseRecordPoolId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "购买记录数组\u003cbr/\u003eId: 标的id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003eBidType: 标计息类型 标记活期定期 \u003cbr/\u003eBidMark: 标的购买限制 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eRemainNumber: 标剩余金额 \u003cbr/\u003eBidNo: 标订单号 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标详情路径 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eCheckDate: 结算日期 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 \u003cbr/\u003eIsVip: 是计入vip的 \u003cbr/\u003eIsRank: 是计入rank的 \u003cbr/\u003eMark: 标记 \u003cbr/\u003eDays: 实际经过的天数 \u003cbr/\u003eBidRate: 标的物利率 \u003cbr/\u003eIncomeDate: 起息时间 \u003cbr/\u003eBidRateBase: 基础利率 \u003cbr/\u003eBidRateRaise: 加息利率 \u003cbr/\u003eNumber: 金额 \u003cbr/\u003eDrawNumber: 持有金额 \u003cbr/\u003eEarningsNumber: 购买金额的收益 \u003cbr/\u003eYesterEarningsNumber: 昨日收益 \u003cbr/\u003eExpireEarningsNumber: 预计到期收益 \u003cbr/\u003eExpireDays: 预计天数 废弃 \u003cbr/\u003ePayNumber: 支付金额 \u003cbr/\u003ePurchaseRecordId: 购买记录id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eIsIncome: 是否起息 \u003cbr/\u003eCouponMsg: 优惠券使用描述 \u003cbr/\u003eMinNumber: 起投金额 \u003cbr/\u003eMaxNumber: 最大金额 \u003cbr/\u003eStepNumber: 增幅金额 \u003cbr/\u003eMaxBuySize: 最大购买笔数 \u003cbr/\u003eRefillDaysLimit: 标的物最大期限 \u003cbr/\u003eRefillBidType: 标计息类型 \u003cbr/\u003eRefillBidMark: 标的购买限制 \u003cbr/\u003eRefillRemainNumberMin: 全部符合条件的标 累计小于最小续标金额 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespPurchaseRecord"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/record/getpurchaseexistbyuserid": {
            "post": {
                "tags": [
                    "/record"
                ],
                "summary": "获取用户还有余额的购买记录",
                "description": "获取用户还有余额的购买记录",
                "operationId": "/record/getpurchaseexistbyuserid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "购买记录数组\u003cbr/\u003eId: 标的id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003eBidType: 标计息类型 标记活期定期 \u003cbr/\u003eBidMark: 标的购买限制 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eRemainNumber: 标剩余金额 \u003cbr/\u003eBidNo: 标订单号 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标详情路径 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eCheckDate: 结算日期 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 \u003cbr/\u003eIsVip: 是计入vip的 \u003cbr/\u003eIsRank: 是计入rank的 \u003cbr/\u003eMark: 标记 \u003cbr/\u003eDays: 实际经过的天数 \u003cbr/\u003eBidRate: 标的物利率 \u003cbr/\u003eIncomeDate: 起息时间 \u003cbr/\u003eBidRateBase: 基础利率 \u003cbr/\u003eBidRateRaise: 加息利率 \u003cbr/\u003eNumber: 金额 \u003cbr/\u003eDrawNumber: 持有金额 \u003cbr/\u003eEarningsNumber: 购买金额的收益 \u003cbr/\u003eYesterEarningsNumber: 昨日收益 \u003cbr/\u003eExpireEarningsNumber: 预计到期收益 \u003cbr/\u003eExpireDays: 预计天数 废弃 \u003cbr/\u003ePayNumber: 支付金额 \u003cbr/\u003ePurchaseRecordId: 购买记录id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eIsIncome: 是否起息 \u003cbr/\u003eCouponMsg: 优惠券使用描述 \u003cbr/\u003eMinNumber: 起投金额 \u003cbr/\u003eMaxNumber: 最大金额 \u003cbr/\u003eStepNumber: 增幅金额 \u003cbr/\u003eMaxBuySize: 最大购买笔数 \u003cbr/\u003eRefillDaysLimit: 标的物最大期限 \u003cbr/\u003eRefillBidType: 标计息类型 \u003cbr/\u003eRefillBidMark: 标的购买限制 \u003cbr/\u003eRefillRemainNumberMin: 全部符合条件的标 累计小于最小续标金额 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespPurchaseRecord"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/record/getpurchasenonexistbyuserid": {
            "post": {
                "tags": [
                    "/record"
                ],
                "summary": "获取用户还有余额的购买记录",
                "description": "获取用户还有余额的购买记录",
                "operationId": "/record/getpurchasenonexistbyuserid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserIdOffsetLimit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "购买记录数组\u003cbr/\u003eId: 标的id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003eBidType: 标计息类型 标记活期定期 \u003cbr/\u003eBidMark: 标的购买限制 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eRemainNumber: 标剩余金额 \u003cbr/\u003eBidNo: 标订单号 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标详情路径 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eCheckDate: 结算日期 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 \u003cbr/\u003eIsVip: 是计入vip的 \u003cbr/\u003eIsRank: 是计入rank的 \u003cbr/\u003eMark: 标记 \u003cbr/\u003eDays: 实际经过的天数 \u003cbr/\u003eBidRate: 标的物利率 \u003cbr/\u003eIncomeDate: 起息时间 \u003cbr/\u003eBidRateBase: 基础利率 \u003cbr/\u003eBidRateRaise: 加息利率 \u003cbr/\u003eNumber: 金额 \u003cbr/\u003eDrawNumber: 持有金额 \u003cbr/\u003eEarningsNumber: 购买金额的收益 \u003cbr/\u003eYesterEarningsNumber: 昨日收益 \u003cbr/\u003eExpireEarningsNumber: 预计到期收益 \u003cbr/\u003eExpireDays: 预计天数 废弃 \u003cbr/\u003ePayNumber: 支付金额 \u003cbr/\u003ePurchaseRecordId: 购买记录id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eIsIncome: 是否起息 \u003cbr/\u003eCouponMsg: 优惠券使用描述 \u003cbr/\u003eMinNumber: 起投金额 \u003cbr/\u003eMaxNumber: 最大金额 \u003cbr/\u003eStepNumber: 增幅金额 \u003cbr/\u003eMaxBuySize: 最大购买笔数 \u003cbr/\u003eRefillDaysLimit: 标的物最大期限 \u003cbr/\u003eRefillBidType: 标计息类型 \u003cbr/\u003eRefillBidMark: 标的购买限制 \u003cbr/\u003eRefillRemainNumberMin: 全部符合条件的标 累计小于最小续标金额 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespPurchaseRecord"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/record/getpurchasepool": {
            "post": {
                "tags": [
                    "/record"
                ],
                "summary": "获取用户购买池记录",
                "description": "获取用户购买池记录",
                "operationId": "/record/getpurchasepool",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"购买池记录id\"\u003cbr/\u003ePurchaseRecordPoolId: ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqPurchaseRecordPoolId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "购买池记录数组\u003cbr/\u003eId: 购买池id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eNumber: 购买金额 \u003cbr/\u003ePayNumber: 支付金额 \u003cbr/\u003eRemainNumber: 持有金额 \u003cbr/\u003eStatus: 持有状态 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.PurchaseRecordPool"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/record/getpurchasepoolbyuserid": {
            "post": {
                "tags": [
                    "/record"
                ],
                "summary": "获取用户的购买池购买记录",
                "description": "获取用户的购买池购买记录",
                "operationId": "/record/getpurchasepoolbyuserid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id 和 索引\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserIdOffsetLimit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "购买记录数组\u003cbr/\u003eId: 购买池id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eNumber: 购买金额 \u003cbr/\u003ePayNumber: 支付金额 \u003cbr/\u003eRemainNumber: 持有金额 \u003cbr/\u003eStatus: 持有状态 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.PurchaseRecordPool"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/record/getpurchasepoolexistbyuserid": {
            "post": {
                "tags": [
                    "/record"
                ],
                "summary": "获取用户还有余额的购买池记录",
                "description": "获取用户还有余额的购买池记录",
                "operationId": "/record/getpurchasepoolexistbyuserid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "购买池记录数组\u003cbr/\u003eId: 购买池id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eNumber: 购买金额 \u003cbr/\u003ePayNumber: 支付金额 \u003cbr/\u003eRemainNumber: 持有金额 \u003cbr/\u003eStatus: 持有状态 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.PurchaseRecordPool"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/record/getpurchaserecordsexistbalance": {
            "post": {
                "tags": [
                    "/record"
                ],
                "summary": "获取标的持有人",
                "description": "获取标的持有人",
                "operationId": "/record/getpurchaserecordsexistbalance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"标的id\"\u003cbr/\u003eBidId: 标的id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBidId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "标的持有的人购买记录\u003cbr/\u003eId: 标的id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003eBidType: 标计息类型 标记活期定期 \u003cbr/\u003eBidMark: 标的购买限制 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eRemainNumber: 标剩余金额 \u003cbr/\u003eBidNo: 标订单号 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标详情路径 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eCheckDate: 结算日期 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 \u003cbr/\u003eIsVip: 是计入vip的 \u003cbr/\u003eIsRank: 是计入rank的 \u003cbr/\u003eMark: 标记 \u003cbr/\u003eDays: 实际经过的天数 \u003cbr/\u003eBidRate: 标的物利率 \u003cbr/\u003eIncomeDate: 起息时间 \u003cbr/\u003eBidRateBase: 基础利率 \u003cbr/\u003eBidRateRaise: 加息利率 \u003cbr/\u003eNumber: 金额 \u003cbr/\u003eDrawNumber: 持有金额 \u003cbr/\u003eEarningsNumber: 购买金额的收益 \u003cbr/\u003eYesterEarningsNumber: 昨日收益 \u003cbr/\u003eExpireEarningsNumber: 预计到期收益 \u003cbr/\u003eExpireDays: 预计天数 废弃 \u003cbr/\u003ePayNumber: 支付金额 \u003cbr/\u003ePurchaseRecordId: 购买记录id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eIsIncome: 是否起息 \u003cbr/\u003eCouponMsg: 优惠券使用描述 \u003cbr/\u003eMinNumber: 起投金额 \u003cbr/\u003eMaxNumber: 最大金额 \u003cbr/\u003eStepNumber: 增幅金额 \u003cbr/\u003eMaxBuySize: 最大购买笔数 \u003cbr/\u003eRefillDaysLimit: 标的物最大期限 \u003cbr/\u003eRefillBidType: 标计息类型 \u003cbr/\u003eRefillBidMark: 标的购买限制 \u003cbr/\u003eRefillRemainNumberMin: 全部符合条件的标 累计小于最小续标金额 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespPurchaseRecord"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/record/getrecordhistory": {
            "post": {
                "tags": [
                    "/record"
                ],
                "summary": "获取购买记录的历史收益记录",
                "description": "获取购买记录的历史收益记录",
                "operationId": "/record/getrecordhistory",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"购买记录id\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetRecordHistory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "购买记录历史数据\u003cbr/\u003e",
                        "schema": {
                            "$ref": "#/definitions/models.HistoryRecords"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/record/getrecordhistorypool": {
            "post": {
                "tags": [
                    "/record"
                ],
                "summary": "获取购买记录池的历史收益记录",
                "description": "获取购买记录池的历史收益记录",
                "operationId": "/record/getrecordhistorypool",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"购买记录id\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetRecordHistory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "购买记录历史数据\u003cbr/\u003e",
                        "schema": {
                            "$ref": "#/definitions/models.HistoryRecords"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/record/getreturned": {
            "post": {
                "tags": [
                    "/record"
                ],
                "summary": "回款的金额和收益",
                "description": "回款的金额和收益",
                "operationId": "/record/getreturned",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003ePurchaseRecordId: 购买记录id \u003cbr/\u003eState: 状态 \u003cbr/\u003eBeginDate: 开始日期 \u003cbr/\u003eEndDate: 结束日期 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetReturned"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eId: \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003ePurchaseRecordId: 购买记录id \u003cbr/\u003eState: 状态 \u003cbr/\u003eStages: 第几期 \u003cbr/\u003eTotalStages: 总的几期 \u003cbr/\u003eDrawNumber: 还款本金 \u003cbr/\u003eEarningsNumber: 还款利息 \u003cbr/\u003eCheckDate: 还款日期 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.Returned"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/record/getreturnedpool": {
            "post": {
                "tags": [
                    "/record"
                ],
                "summary": "获取购买池赎回记录",
                "description": "获取购买池赎回记录",
                "operationId": "/record/getreturnedpool",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"索引\"\u003cbr/\u003eBeginTime: 起始日期 \u003cbr/\u003eEndTime: 结束日期 \u003cbr/\u003eOffset: 起始偏移 \u003cbr/\u003eLimit: 限制取多少条 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqDateOffsetLimit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "购买记录数组\u003cbr/\u003eId: \u003cbr/\u003eUserId: 用户 id \u003cbr/\u003eBidId: 标 id \u003cbr/\u003ePurchaseRecordId: 购买记录id \u003cbr/\u003ePurchaseRecordPoolId: 购买记录池id \u003cbr/\u003eDays: 存款天数 \u003cbr/\u003eDrawNumber: 回款本金 \u003cbr/\u003eEarningsNumber: 回款利息 \u003cbr/\u003eYesterEarningsNumber: 昨天回款利息 \u003cbr/\u003eCreateTime: ",
                        "schema": {
                            "$ref": "#/definitions/models.ReturnedRecord"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/record/getreturnedpoolbyuserid": {
            "post": {
                "tags": [
                    "/record"
                ],
                "summary": "获取用户的购买池赎回记录",
                "description": "获取用户的购买池赎回记录",
                "operationId": "/record/getreturnedpoolbyuserid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id 和 索引\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserIdOffsetLimit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "购买记录数组\u003cbr/\u003eId: \u003cbr/\u003eUserId: 用户 id \u003cbr/\u003eBidId: 标 id \u003cbr/\u003ePurchaseRecordId: 购买记录id \u003cbr/\u003ePurchaseRecordPoolId: 购买记录池id \u003cbr/\u003eDays: 存款天数 \u003cbr/\u003eDrawNumber: 回款本金 \u003cbr/\u003eEarningsNumber: 回款利息 \u003cbr/\u003eYesterEarningsNumber: 昨天回款利息 \u003cbr/\u003eCreateTime: ",
                        "schema": {
                            "$ref": "#/definitions/models.ReturnedRecord"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/record/getreturnedsum": {
            "post": {
                "tags": [
                    "/record"
                ],
                "summary": "已经回款的金额和收益",
                "description": "已经回款的金额和收益",
                "operationId": "/record/getreturnedsum",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eDrawNumber: 历史投资金额 \u003cbr/\u003eEarningsNumber: 历史投资收益 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespReturnedSum"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/record/getuserreturnrecord": {
            "post": {
                "tags": [
                    "/record"
                ],
                "summary": "已经回款的金额和收益",
                "description": "已经回款的金额和收益",
                "operationId": "/record/getuserreturnrecord",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eDrawNumber: 历史投资金额 \u003cbr/\u003eEarningsNumber: 历史投资收益 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespReturnedSum"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/statistics/chanum": {
            "post": {
                "tags": [
                    "/statistics"
                ],
                "summary": "每日充值统计",
                "description": "每日充值统计",
                "operationId": "/statistics/chanum",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"日期\"\u003cbr/\u003eBeginTime: 起始日期 \u003cbr/\u003eEndTime: 结束日期 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBeginEndDate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eChargeNumber: \u003cbr/\u003eDateTime: ",
                        "schema": {
                            "$ref": "#/definitions/models.ResqStatisticsChargeInfo"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/statistics/countpurchaserecordspooltoday": {
            "post": {
                "tags": [
                    "/statistics"
                ],
                "summary": "获取活期某天收益总和",
                "description": "获取活期某天收益总和",
                "operationId": "/statistics/countpurchaserecordspooltoday",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"日期\"\u003cbr/\u003eBeginTime: 起始日期 \u003cbr/\u003eEndTime: 结束日期 \u003cbr/\u003eOffset: 起始偏移 \u003cbr/\u003eLimit: 限制取多少条 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqDateOffsetLimit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "收益总和\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/statistics/countpurchaserecordstoday": {
            "post": {
                "tags": [
                    "/statistics"
                ],
                "summary": "获取定期购买记录数量",
                "description": "获取定期购买记录数量",
                "operationId": "/statistics/countpurchaserecordstoday",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"日期\"\u003cbr/\u003eBeginTime: 起始日期 \u003cbr/\u003eEndTime: 结束日期 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBeginEndDate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "收益总和\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/statistics/getchargerecordtoday": {
            "post": {
                "tags": [
                    "/statistics"
                ],
                "summary": "获取某天充值记录",
                "description": "获取某天充值记录",
                "operationId": "/statistics/getchargerecordtoday",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"日期\"\u003cbr/\u003eBeginTime: 起始日期 \u003cbr/\u003eEndTime: 结束日期 \u003cbr/\u003eOffset: 起始偏移 \u003cbr/\u003eLimit: 限制取多少条 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqDateOffsetLimit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "收益总和\u003cbr/\u003eId: 充值记录id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eNumber: 支付金额 \u003cbr/\u003ePayWay: 支付方式 \u003cbr/\u003eBankcardId: 银行卡id \u003cbr/\u003ePayNo: 订单号 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.ChargeRecord"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/statistics/getestcheckbids": {
            "post": {
                "tags": [
                    "/statistics"
                ],
                "summary": "获取某天回款数据 用getestcheckreturned",
                "description": "获取某天回款数据",
                "operationId": "/statistics/getestcheckbids",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"日期\"\u003cbr/\u003eDate: 日期 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqDate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "收益总和\u003cbr/\u003eId: 废除 \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBidName: 标名 \u003cbr/\u003eBidId: 标id \u003cbr/\u003eDaysLimit: 标的物总天数 \u003cbr/\u003ePurchaseRecordId: 购买记录id \u003cbr/\u003ePurchaseRecordPoolId: 购买记录池id \u003cbr/\u003eState: 状态 \u003cbr/\u003eStages: 第几期 \u003cbr/\u003eTotalStages: 共几期 \u003cbr/\u003eDays: 存款天数 废除 \u003cbr/\u003eDrawNumber: 回款本金 \u003cbr/\u003eEarningsNumber: 回款利息 \u003cbr/\u003eYesterEarningsNumber: 昨天回款利息 废除 \u003cbr/\u003eCheckDate: 还款日期 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespReturnedRecordInfo"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/statistics/getestcheckreturned": {
            "post": {
                "tags": [
                    "/statistics"
                ],
                "summary": "获取某天回款数据 2017-07-18",
                "description": "获取某天回款数据",
                "operationId": "/statistics/getestcheckreturned",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"日期\"\u003cbr/\u003eBeginTime: 起始日期 可选 \u003cbr/\u003eEndTime: 结束日期 可选 \u003cbr/\u003eUserId: 用户id \u003cbr/\u003ePurchaseRecordId: 购买记录 id \u003cbr/\u003eStatus: 是否已经回款 可选 0 | ReturnedStateBefore | ReturnedStateAfter \u003cbr/\u003eReal: 真实数据 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetEstCheckReturned"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "收益总和\u003cbr/\u003eId: 废除 \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBidName: 标名 \u003cbr/\u003eBidId: 标id \u003cbr/\u003eDaysLimit: 标的物总天数 \u003cbr/\u003ePurchaseRecordId: 购买记录id \u003cbr/\u003ePurchaseRecordPoolId: 购买记录池id \u003cbr/\u003eState: 状态 \u003cbr/\u003eStages: 第几期 \u003cbr/\u003eTotalStages: 共几期 \u003cbr/\u003eDays: 存款天数 废除 \u003cbr/\u003eDrawNumber: 回款本金 \u003cbr/\u003eEarningsNumber: 回款利息 \u003cbr/\u003eYesterEarningsNumber: 昨天回款利息 废除 \u003cbr/\u003eCheckDate: 还款日期 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespReturnedRecordInfo"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/statistics/getpurchaserecordspooltoday": {
            "post": {
                "tags": [
                    "/statistics"
                ],
                "summary": "获取活期某天收益总和",
                "description": "获取活期某天收益总和",
                "operationId": "/statistics/getpurchaserecordspooltoday",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"日期\"\u003cbr/\u003eBeginTime: 起始日期 \u003cbr/\u003eEndTime: 结束日期 \u003cbr/\u003eOffset: 起始偏移 \u003cbr/\u003eLimit: 限制取多少条 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqDateOffsetLimit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "收益总和\u003cbr/\u003eId: 购买池id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eNumber: 购买金额 \u003cbr/\u003ePayNumber: 支付金额 \u003cbr/\u003eRemainNumber: 持有金额 \u003cbr/\u003eStatus: 持有状态 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.PurchaseRecordPool"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/statistics/getpurchaserecordspooltodaysum": {
            "post": {
                "tags": [
                    "/statistics"
                ],
                "summary": "获取活期某天收益总和",
                "description": "获取活期某天收益总和",
                "operationId": "/statistics/getpurchaserecordspooltodaysum",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"日期\"\u003cbr/\u003eBeginTime: 起始日期 \u003cbr/\u003eEndTime: 结束日期 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBeginEndDate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "收益总和\u003cbr/\u003eNumber: 购买金额 \u003cbr/\u003ePayNumber: 支付金额 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespToday"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/statistics/getpurchaserecordstoday": {
            "post": {
                "tags": [
                    "/statistics"
                ],
                "summary": "获取定期购买记录",
                "description": "获取定期购买记录",
                "operationId": "/statistics/getpurchaserecordstoday",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"日期\"\u003cbr/\u003eBeginTime: 起始日期 \u003cbr/\u003eEndTime: 结束日期 \u003cbr/\u003eOffset: 起始偏移 \u003cbr/\u003eLimit: 限制取多少条 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqDateOffsetLimit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "收益总和\u003cbr/\u003eId: 标的id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003eBidType: 标计息类型 标记活期定期 \u003cbr/\u003eBidMark: 标的购买限制 \u003cbr/\u003eDaysLimit: 这是多少天标 \u003cbr/\u003eTotalNumber: 标总金额 \u003cbr/\u003eRemainNumber: 标剩余金额 \u003cbr/\u003eBidNo: 标订单号 \u003cbr/\u003eBidDesc: 标描述 \u003cbr/\u003eBidName: 标名称 \u003cbr/\u003eBidInfoPath: 标详情路径 \u003cbr/\u003eBidRateShow: 标利率特殊显示 \u003cbr/\u003eCheckDate: 结算日期 \u003cbr/\u003eBeginTime: 开卖时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 \u003cbr/\u003eIsVip: 是计入vip的 \u003cbr/\u003eIsRank: 是计入rank的 \u003cbr/\u003eMark: 标记 \u003cbr/\u003eDays: 实际经过的天数 \u003cbr/\u003eBidRate: 标的物利率 \u003cbr/\u003eIncomeDate: 起息时间 \u003cbr/\u003eBidRateBase: 基础利率 \u003cbr/\u003eBidRateRaise: 加息利率 \u003cbr/\u003eNumber: 金额 \u003cbr/\u003eDrawNumber: 持有金额 \u003cbr/\u003eEarningsNumber: 购买金额的收益 \u003cbr/\u003eYesterEarningsNumber: 昨日收益 \u003cbr/\u003eExpireEarningsNumber: 预计到期收益 \u003cbr/\u003eExpireDays: 预计天数 废弃 \u003cbr/\u003ePayNumber: 支付金额 \u003cbr/\u003ePurchaseRecordId: 购买记录id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eIsIncome: 是否起息 \u003cbr/\u003eCouponMsg: 优惠券使用描述 \u003cbr/\u003eMinNumber: 起投金额 \u003cbr/\u003eMaxNumber: 最大金额 \u003cbr/\u003eStepNumber: 增幅金额 \u003cbr/\u003eMaxBuySize: 最大购买笔数 \u003cbr/\u003eRefillDaysLimit: 标的物最大期限 \u003cbr/\u003eRefillBidType: 标计息类型 \u003cbr/\u003eRefillBidMark: 标的购买限制 \u003cbr/\u003eRefillRemainNumberMin: 全部符合条件的标 累计小于最小续标金额 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespPurchaseRecord"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/statistics/getpurchaserecordstodaysum": {
            "post": {
                "tags": [
                    "/statistics"
                ],
                "summary": "获取定期某天收益总和",
                "description": "获取定期某天收益总和",
                "operationId": "/statistics/getpurchaserecordstodaysum",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"日期\"\u003cbr/\u003eBeginTime: 起始日期 \u003cbr/\u003eEndTime: 结束日期 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBeginEndDate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "收益总和\u003cbr/\u003eNumber: 购买金额 \u003cbr/\u003ePayNumber: 支付金额 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespToday"
                        }
                    },
                    "202": {
                        "description": "参数错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/statistics/purnum": {
            "post": {
                "tags": [
                    "/statistics"
                ],
                "summary": "每日购买统计",
                "description": "每日购买统计",
                "operationId": "/statistics/purnum",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"日期\"\u003cbr/\u003eBeginTime: 起始日期 \u003cbr/\u003eEndTime: 结束日期 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBeginEndDate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eNumber: \u003cbr/\u003eDiscountNumber: \u003cbr/\u003eDateTime: ",
                        "schema": {
                            "$ref": "#/definitions/models.ResqStatisticsPurchaseInfo"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/statistics/regnum": {
            "post": {
                "tags": [
                    "/statistics"
                ],
                "summary": "每日注册人数统计",
                "description": "每日注册人数统计",
                "operationId": "/statistics/regnum",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"日期\"\u003cbr/\u003eBeginTime: 起始日期 \u003cbr/\u003eEndTime: 结束日期 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBeginEndDate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eRegisterNumber: \u003cbr/\u003eDateTime: ",
                        "schema": {
                            "$ref": "#/definitions/models.ResqStatisticsRegisterInfo"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/statistics/retnum": {
            "post": {
                "tags": [
                    "/statistics"
                ],
                "summary": "每日回款统计",
                "description": "每日回款统计",
                "operationId": "/statistics/retnum",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"日期\"\u003cbr/\u003eBeginTime: 起始日期 \u003cbr/\u003eEndTime: 结束日期 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBeginEndDate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eDrawNumber: \u003cbr/\u003eEarningsNumber: \u003cbr/\u003eDateTime: ",
                        "schema": {
                            "$ref": "#/definitions/models.ResqStatisticsReturnInfo"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/statistics/returningearns": {
            "post": {
                "tags": [
                    "/statistics"
                ],
                "summary": "根据标记录获取投资金额",
                "description": "每日充值统计",
                "operationId": "/statistics/returningearns",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"日期\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBidState: 募集状态 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserBidState"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eChargeNumber: \u003cbr/\u003eDateTime: ",
                        "schema": {
                            "$ref": "#/definitions/models.ResqStatisticsChargeInfo"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/statistics/returningnum": {
            "post": {
                "tags": [
                    "/statistics"
                ],
                "summary": "根据未回款标的投资",
                "description": "每日充值统计",
                "operationId": "/statistics/returningnum",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"日期\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBidState: 募集状态 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserBidState"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eChargeNumber: \u003cbr/\u003eDateTime: ",
                        "schema": {
                            "$ref": "#/definitions/models.ResqStatisticsChargeInfo"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/statistics/vernum": {
            "post": {
                "tags": [
                    "/statistics"
                ],
                "summary": "每日认证统计",
                "description": "每日认证统计",
                "operationId": "/statistics/vernum",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"日期\"\u003cbr/\u003eBeginTime: 起始日期 \u003cbr/\u003eEndTime: 结束日期 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBeginEndDate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eVerifyNumber: \u003cbr/\u003eDateTime: ",
                        "schema": {
                            "$ref": "#/definitions/models.ResqStatisticsVerifyInfo"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/task/changebidtype": {
            "post": {
                "tags": [
                    "/task"
                ],
                "summary": "改变老标 已满标的类型",
                "description": "改变老标 已满标的类型 请在回款后执行",
                "operationId": "/task/changebidtype",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eBidId: 标的id \u003cbr/\u003eBidType: 标的类型 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqChangeBidType"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/task/checkmatchbids": {
            "post": {
                "tags": [
                    "/task"
                ],
                "summary": "回款和活期匹配的定时任务",
                "description": "回款和活期匹配的定时任务",
                "operationId": "/task/checkmatchbids",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "执行信息描述\u003cbr/\u003e"
                    }
                }
            }
        },
        "/task/syncrefill": {
            "post": {
                "tags": [
                    "/task"
                ],
                "summary": "尝试续标 10s内多次请求忽略",
                "description": "尝试续标 10s内多次请求忽略",
                "operationId": "/task/syncrefill",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "续标成功\u003cbr/\u003e"
                    }
                }
            }
        },
        "/transaction/addpurchase": {
            "post": {
                "tags": [
                    "/transaction"
                ],
                "summary": "添加用户的购买记录 这个接口是一个完整的购买流程 这里不验证 标的购买限制",
                "description": "添加用户的购买记录 这个接口是一个完整的购买流程 这里不验证 标的购买限制",
                "operationId": "/transaction/addpurchase",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id 标id 购买金额 必须\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003eNumber: 购买金额 \u003cbr/\u003eCouponRules: 购买时使用的券 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqAddPurchase"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "购买成功返回购买记录id\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "购买失败有各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/transaction/addpurchasepool": {
            "post": {
                "tags": [
                    "/transaction"
                ],
                "summary": "添加用户的购买池记录 这个接口是一个完整的购买流程 这里不验证 标的购买限制",
                "description": "添加用户的购买池记录 这个接口是一个完整的购买流程 这里不验证 标的购买限制",
                "operationId": "/transaction/addpurchasepool",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id 购买金额 必须\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eNumber: 购买金额 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqAddPurchasePool"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "购买成功返回购买记录id\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "购买失败有各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/transaction/addreturned": {
            "post": {
                "tags": [
                    "/transaction"
                ],
                "summary": "添加赎回记录 这个接口是一个完整的赎回流程 这里不验证 标的赎回限制",
                "description": "添加赎回记录 这个接口是一个完整的赎回流程 这里不验证 标的赎回限制",
                "operationId": "/transaction/addreturned",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"购买记录id\"\u003cbr/\u003ePurchaseRecordId: 购买记录id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqAddReturned"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "赎回成功返回回款记录id\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "赎回失败有各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/transaction/addreturnedpool": {
            "post": {
                "tags": [
                    "/transaction"
                ],
                "summary": "添加赎回记录 这个接口是一个完整的赎回流程 这里不验证 标的赎回限制",
                "description": "添加赎回记录 这个接口是一个完整的赎回流程 这里不验证 标的赎回限制",
                "operationId": "/transaction/addreturnedpool",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"购买池记录id\"\u003cbr/\u003ePurchaseRecordPoolId: 购买记录id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqAddReturnedPool"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "赎回成功返回回款记录id\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "赎回失败有各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/transaction/estaddpurchase": {
            "post": {
                "tags": [
                    "/transaction"
                ],
                "summary": "添加用户的购买记录 AddPurchase接口的预览不作出实际改变 2017-09-17",
                "description": "添加用户的购买记录 AddPurchase接口的预览不作出实际改变",
                "operationId": "/transaction/estaddpurchase",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id 标id 购买金额 必须\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003eNumber: 购买金额 \u003cbr/\u003eCouponRules: 购买时使用的券 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqAddPurchase"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "模拟赎回成功\u003cbr/\u003eId: \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003ePurchaseRecordPoolId: 购买记录池id \u003cbr/\u003eNumber: 购买金额 \u003cbr/\u003ePayNumber: 支付金额 \u003cbr/\u003eRemainNumber: 剩余金额 \u003cbr/\u003eCouponRules: 购买时使用的券 \u003cbr/\u003eEarningsNumber: 预期收益 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespEstAddPurchase"
                        }
                    },
                    "202": {
                        "description": "购买失败有各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/transaction/estaddpurchasepool": {
            "post": {
                "tags": [
                    "/transaction"
                ],
                "summary": "添加用户的购买池记录 AddPurchasePool接口的预览不作出实际改变",
                "description": "添加用户的购买池记录 AddPurchasePool接口的预览不作出实际改变",
                "operationId": "/transaction/estaddpurchasepool",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id 购买金额 必须\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eNumber: 购买金额 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqAddPurchasePool"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "模拟赎回成功\u003cbr/\u003eId: 购买池id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eNumber: 购买金额 \u003cbr/\u003ePayNumber: 支付金额 \u003cbr/\u003eRemainNumber: 持有金额 \u003cbr/\u003eStatus: 持有状态 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.PurchaseRecordPool"
                        }
                    },
                    "202": {
                        "description": "购买失败有各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/transaction/estaddreturned": {
            "post": {
                "tags": [
                    "/transaction"
                ],
                "summary": "添加赎回记录  AddReturned接口的预览不作出实际改变 废弃 用EstAddReturnedRecordNew",
                "description": "添加赎回记录  AddReturned接口的预览不作出实际改变",
                "operationId": "/transaction/estaddreturned",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"购买记录id\"\u003cbr/\u003ePurchaseRecordId: 购买记录id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqAddReturned"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "模拟赎回成功\u003cbr/\u003eId: 废除 \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBidName: 标名 \u003cbr/\u003eBidId: 标id \u003cbr/\u003eDaysLimit: 标的物总天数 \u003cbr/\u003ePurchaseRecordId: 购买记录id \u003cbr/\u003ePurchaseRecordPoolId: 购买记录池id \u003cbr/\u003eState: 状态 \u003cbr/\u003eStages: 第几期 \u003cbr/\u003eTotalStages: 共几期 \u003cbr/\u003eDays: 存款天数 废除 \u003cbr/\u003eDrawNumber: 回款本金 \u003cbr/\u003eEarningsNumber: 回款利息 \u003cbr/\u003eYesterEarningsNumber: 昨天回款利息 废除 \u003cbr/\u003eCheckDate: 还款日期 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespReturnedRecordInfo"
                        }
                    },
                    "202": {
                        "description": "赎回失败有各种原因\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/transaction/estaddreturnedpool": {
            "post": {
                "tags": [
                    "/transaction"
                ],
                "summary": "添加赎回记录 这个接口是一个完整的赎回流程 这里不验证 标的赎回限制 尝试",
                "description": "添加赎回记录 这个接口是一个完整的赎回流程 这里不验证 标的赎回限制 尝试",
                "operationId": "/transaction/estaddreturnedpool",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"购买池记录id\"\u003cbr/\u003ePurchaseRecordPoolId: 购买记录id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqAddReturnedPool"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "模拟赎回成功\u003cbr/\u003eId: \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003ePurchaseRecordPoolId: 购买记录池id \u003cbr/\u003eNumber: 购买金额 \u003cbr/\u003ePayNumber: 支付金额 \u003cbr/\u003eRemainNumber: 剩余金额 \u003cbr/\u003eCouponRules: 购买时使用的券 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.PurchaseRecord"
                        }
                    },
                    "202": {
                        "description": "赎回失败有各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/transaction/estaddreturnedrecordnew": {
            "post": {
                "tags": [
                    "/transaction"
                ],
                "summary": "预计回款计划表",
                "description": "预计回款计划表",
                "operationId": "/transaction/estaddreturnedrecordnew",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"购买记录id\"\u003cbr/\u003ePurchaseRecordId: 购买记录id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqAddReturned"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "模拟赎回成功\u003cbr/\u003eId: 废除 \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBidName: 标名 \u003cbr/\u003eBidId: 标id \u003cbr/\u003eDaysLimit: 标的物总天数 \u003cbr/\u003ePurchaseRecordId: 购买记录id \u003cbr/\u003ePurchaseRecordPoolId: 购买记录池id \u003cbr/\u003eState: 状态 \u003cbr/\u003eStages: 第几期 \u003cbr/\u003eTotalStages: 共几期 \u003cbr/\u003eDays: 存款天数 废除 \u003cbr/\u003eDrawNumber: 回款本金 \u003cbr/\u003eEarningsNumber: 回款利息 \u003cbr/\u003eYesterEarningsNumber: 昨天回款利息 废除 \u003cbr/\u003eCheckDate: 还款日期 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespReturnedRecordInfo"
                        }
                    },
                    "202": {
                        "description": "赎回失败有各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/usercoupon/addusercoupon": {
            "post": {
                "tags": [
                    "/usercoupon"
                ],
                "summary": "添加用户优惠券",
                "description": "添加用户优惠券",
                "operationId": "/usercoupon/addusercoupon",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eCouponId: 优惠券类型 id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqAddUserCoupon"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回优惠券数组\u003cbr/\u003eId: 用户优惠券id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eCouponId: 优惠券id \u003cbr/\u003eState: 使用状态 \u003cbr/\u003eBeginTime: 生效时间 \u003cbr/\u003eEndTime: 失效时间 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.UserCoupon"
                        }
                    },
                    "202": {
                        "description": "操作失败各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/usercoupon/delusercoupon": {
            "post": {
                "tags": [
                    "/usercoupon"
                ],
                "summary": "设置优惠券失效",
                "description": "设置优惠券失效",
                "operationId": "/usercoupon/delusercoupon",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"优惠券类型id\"\u003cbr/\u003eUserCouponId: 用户的优惠券id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserCouponId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "操作成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "操作失败各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/usercoupon/getallcoupon": {
            "post": {
                "tags": [
                    "/usercoupon"
                ],
                "summary": "获取优惠券类型",
                "description": "获取优惠券类型",
                "operationId": "/usercoupon/getallcoupon",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "优惠券类型\u003cbr/\u003eId: 优惠券类型id \u003cbr/\u003eCouponType: 优惠券类型  0 | CouponTypeDiscount | CouponTypeRaise \u003cbr/\u003eCouponRule: 优惠券规则 \u003cbr/\u003eCouponName: 优惠券描述 \u003cbr/\u003eCouponDesc: 优惠券信息 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.Coupon"
                        }
                    },
                    "202": {
                        "description": "操作失败各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/usercoupon/getusercouponbyuserid": {
            "post": {
                "tags": [
                    "/usercoupon"
                ],
                "summary": "获取用户可以使用的优惠券 全部",
                "description": "获取用户可以使用的优惠券 全部",
                "operationId": "/usercoupon/getusercouponbyuserid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBeginTime: 起始时间 \u003cbr/\u003eEndTime: 截止时间 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetUserCoupon"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回优惠券数组\u003cbr/\u003eId: 用户优惠券id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eCouponId: 优惠券id \u003cbr/\u003eState: 使用状态 \u003cbr/\u003eBeginTime: 生效时间 \u003cbr/\u003eEndTime: 失效时间 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.UserCoupon"
                        }
                    },
                    "202": {
                        "description": "操作失败各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/usercoupon/getusercouponcanuse": {
            "post": {
                "tags": [
                    "/usercoupon"
                ],
                "summary": "获取用户买某一笔标可以使用的优惠券",
                "description": "获取用户买某一笔标可以使用的优惠券",
                "operationId": "/usercoupon/getusercouponcanuse",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003eNumber: 金额 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetUserCouponCanUse"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回优惠券数组\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eBidId: 标的id \u003cbr/\u003eNumber: 金额 ",
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetUserCouponCanUse"
                        }
                    },
                    "202": {
                        "description": "操作失败各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/usercoupon/getusercouponcanusebyuserid": {
            "post": {
                "tags": [
                    "/usercoupon"
                ],
                "summary": "获取用户的能使用的优惠券 分页 2017-07-03",
                "description": "获取用户的能使用的优惠券 分页",
                "operationId": "/usercoupon/getusercouponcanusebyuserid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserIdOffsetLimit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回优惠券数组\u003cbr/\u003eId: 用户优惠券id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eCouponId: 优惠券id \u003cbr/\u003eState: 使用状态 \u003cbr/\u003eBeginTime: 生效时间 \u003cbr/\u003eEndTime: 失效时间 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.UserCoupon"
                        }
                    },
                    "202": {
                        "description": "操作失败各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/usercoupon/getusercouponnew": {
            "post": {
                "tags": [
                    "/usercoupon"
                ],
                "summary": "获取用户可以使用的优惠券 筛选 2017-09-11",
                "description": "获取用户可以使用的优惠券 筛选",
                "operationId": "/usercoupon/getusercouponnew",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserId: 用户id 可选 \u003cbr/\u003eState: 优惠券状态 可选 0 | UserCouponStateUnused | UserCouponStateUsed | UserCouponStateOverdue \u003cbr/\u003eCouponId: 优惠券id 可选 \u003cbr/\u003eCouponType: 优惠券类型 0 | CouponTypeDiscount | CouponTypeRaise \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetUserCouponNew"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回优惠券数组\u003cbr/\u003eUserCouponId: 优惠券id \u003cbr/\u003eState: 优惠券状态 可选 0 | UserCouponStateUnused | UserCouponStateUsed | UserCouponStateOverdue \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eCouponId: 优惠券类型id \u003cbr/\u003eCouponType: 优惠券类型 0 | CouponTypeDiscount | CouponTypeRaise \u003cbr/\u003eCouponRule: 优惠券规则 \u003cbr/\u003eCouponName: 优惠券描述 \u003cbr/\u003eCouponDesc: 优惠券信息 \u003cbr/\u003eLogo: logo \u003cbr/\u003eBeginTime: 生效时间 \u003cbr/\u003eEndTime: 失效时间 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespGetUserCouponNew"
                        }
                    },
                    "202": {
                        "description": "操作失败各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/usercoupon/getusercouponnotuse": {
            "post": {
                "tags": [
                    "/usercoupon"
                ],
                "summary": "获取用户无法使用的优惠券 全部",
                "description": "获取用户无法使用的优惠券 全部",
                "operationId": "/usercoupon/getusercouponnotuse",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回优惠券数组\u003cbr/\u003eId: 用户优惠券id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eCouponId: 优惠券id \u003cbr/\u003eState: 使用状态 \u003cbr/\u003eBeginTime: 生效时间 \u003cbr/\u003eEndTime: 失效时间 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.UserCoupon"
                        }
                    },
                    "202": {
                        "description": "操作失败各种原因\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/usercoupon/getusercouponoverduebyuserid": {
            "post": {
                "tags": [
                    "/usercoupon"
                ],
                "summary": "获取用户的过期了的优惠券 分页 2017-07-03",
                "description": "获取用户的过期了的优惠券 分页",
                "operationId": "/usercoupon/getusercouponoverduebyuserid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserIdOffsetLimit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回优惠券数组\u003cbr/\u003eId: 用户优惠券id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eCouponId: 优惠券id \u003cbr/\u003eState: 使用状态 \u003cbr/\u003eBeginTime: 生效时间 \u003cbr/\u003eEndTime: 失效时间 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.UserCoupon"
                        }
                    },
                    "202": {
                        "description": "操作失败各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/usercoupon/getusercouponunused": {
            "post": {
                "tags": [
                    "/usercoupon"
                ],
                "summary": "获取用户可以使用的优惠券 全部",
                "description": "获取用户可以使用的优惠券 全部",
                "operationId": "/usercoupon/getusercouponunused",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回优惠券数组\u003cbr/\u003eId: 用户优惠券id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eCouponId: 优惠券id \u003cbr/\u003eState: 使用状态 \u003cbr/\u003eBeginTime: 生效时间 \u003cbr/\u003eEndTime: 失效时间 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.UserCoupon"
                        }
                    },
                    "202": {
                        "description": "操作失败各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/usercoupon/getusercouponused": {
            "post": {
                "tags": [
                    "/usercoupon"
                ],
                "summary": "获取用户已使用的优惠券 全部",
                "description": "获取用户已使用的优惠券 全部",
                "operationId": "/usercoupon/getusercouponused",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回优惠券数组\u003cbr/\u003eId: 用户优惠券id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eCouponId: 优惠券id \u003cbr/\u003eState: 使用状态 \u003cbr/\u003eBeginTime: 生效时间 \u003cbr/\u003eEndTime: 失效时间 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.UserCoupon"
                        }
                    },
                    "202": {
                        "description": "操作失败各种原因\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/usercoupon/getusercouponusedbyuserid": {
            "post": {
                "tags": [
                    "/usercoupon"
                ],
                "summary": "获取用户的使用了的优惠券 分页 2017-07-03",
                "description": "获取用户的使用了的优惠券 分页",
                "operationId": "/usercoupon/getusercouponusedbyuserid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserIdOffsetLimit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回优惠券数组\u003cbr/\u003eId: 用户优惠券id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eCouponId: 优惠券id \u003cbr/\u003eState: 使用状态 \u003cbr/\u003eBeginTime: 生效时间 \u003cbr/\u003eEndTime: 失效时间 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.UserCoupon"
                        }
                    },
                    "202": {
                        "description": "操作失败各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/addbalance": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "添加余额",
                "description": "添加余额",
                "operationId": "/users/addbalance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"添加余额\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eNumber: 金额 \u003cbr/\u003eBankcardId: 银行卡号 \u003cbr/\u003ePayWay: 支付方式类型 \u003cbr/\u003ePayNo: 支付订单号 \u003cbr/\u003eRecordDesc: 来自描述 \u003cbr/\u003eFromId: 来自id \u003cbr/\u003eFromType: 来自类型 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqAddBalance"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "正确\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/addbankcard": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "设置用户身份证与姓名",
                "description": "设置用户身份证与姓名",
                "operationId": "/users/addbankcard",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"交易密码\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eCard: 卡号 \u003cbr/\u003ePhone: 预留手机号 \u003cbr/\u003ePayWay: 支付渠道 \u003cbr/\u003eBindId: 绑定id \u003cbr/\u003eBankCode: 银行编码 \u003cbr/\u003eBankName: 银行名称 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBankcard"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "设置成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "设置失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/addexperience": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "添加用户体验金",
                "description": "添加用户体验金",
                "operationId": "/users/addexperience",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"体验金\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eNumber: 金额 \u003cbr/\u003eMessage: 描述 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqAddExperience"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "体验金\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求出错\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/checkdevicenumber": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "检查设备号",
                "description": "检查设备号",
                "operationId": "/users/checkdevicenumber",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"交易密码\"\u003cbr/\u003eDeviceNumber: 设备号 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqDeviceNumber"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "正确\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/checkpaypwd": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "检查交易密码是否正确",
                "description": "检查交易密码是否正确",
                "operationId": "/users/checkpaypwd",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"交易密码\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003ePwd: 用户密码 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqPwd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "正确\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/countalluser": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "统计所有注册用户",
                "description": "统计所有注册用户",
                "operationId": "/users/countalluser",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/countuserssumbyquery": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "搜索用户",
                "description": "搜索用户",
                "operationId": "/users/countuserssumbyquery",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"查询用户\"\u003cbr/\u003eQuery: 查询 \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetUserByQuery"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登录成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "登录失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/getallusertag": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "获得用户所有的tag",
                "description": "获得用户所有的tag",
                "operationId": "/users/getallusertag",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e",
                        "schema": {
                            "$ref": "#/definitions/map[string]string"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/getbalance": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "获取用户余额",
                "description": "获取用户余额",
                "operationId": "/users/getbalance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"获取用户的余额\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回用户的余额和冻结余额\u003cbr/\u003eId: 用户余额id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eNumber: 余额 \u003cbr/\u003eFreezeNumber: 冻结收益 \u003cbr/\u003eHistoryNormalNumber: 定期历史收益 \u003cbr/\u003eHistoryDemandNumber: 活期历史收益 \u003cbr/\u003eUpdateTime: 更新时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.Balance"
                        }
                    },
                    "202": {
                        "description": "失败各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/getbankcard": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "获取用户银行卡",
                "description": "获取用户银行卡",
                "operationId": "/users/getbankcard",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"获取用户的银行卡\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回用户的余额和冻结余额\u003cbr/\u003eId: \u003cbr/\u003eUserId: \u003cbr/\u003eState: \u003cbr/\u003ePayWay: \u003cbr/\u003eTotalIn: \u003cbr/\u003eTotalOut: \u003cbr/\u003ePhone: \u003cbr/\u003eCard: \u003cbr/\u003eBankCode: \u003cbr/\u003eBindId: \u003cbr/\u003eBankName: \u003cbr/\u003eBankDesc: \u003cbr/\u003eCreateTime: \u003cbr/\u003eSingleLimit: \u003cbr/\u003eDayLimit: ",
                        "schema": {
                            "$ref": "#/definitions/models.RespBankcard"
                        }
                    },
                    "202": {
                        "description": "失败各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/getexperience": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "获取用户体验金列表",
                "description": "获取用户体验金列表",
                "operationId": "/users/getexperience",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "体验金\u003cbr/\u003eId: 体验金id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eNumber: 金额 \u003cbr/\u003eEarningsNumber: 回款金额 \u003cbr/\u003eMessage: 描述 \u003cbr/\u003eStatus: 状态 \u003cbr/\u003eCheckDate: 回款日 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.Experience"
                        }
                    },
                    "202": {
                        "description": "请求出错\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/getexperiencehistory": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "获取用户体验金历史信息",
                "description": "获取用户体验金历史信息",
                "operationId": "/users/getexperiencehistory",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "体验金\u003cbr/\u003eNumber: \u003cbr/\u003eEarningsNumber: ",
                        "schema": {
                            "$ref": "#/definitions/models.RespExperienceHistory"
                        }
                    },
                    "202": {
                        "description": "请求出错\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/getinviteelist": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "获取用户邀请对象列表",
                "description": "获取用户邀请对象列表",
                "operationId": "/users/getinviteelist",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "[]models.RespInviteeList\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求错误\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/getplatformuserrecord": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "渠道用户查询",
                "description": "渠道用户查询",
                "operationId": "/users/getplatformuserrecord",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserName: 实名 \u003cbr/\u003eAccount: 账号 \u003cbr/\u003eBeginTime: 开始时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqPlatformUserTime"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eRegisterNumber: \u003cbr/\u003eDateTime: ",
                        "schema": {
                            "$ref": "#/definitions/models.ResqUserRegisteInfo"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/getpurchaseusertotal": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "投资汇总",
                "description": "一段时间内的投资人数",
                "operationId": "/users/getpurchaseusertotal",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eBeginTime: 起始日期 \u003cbr/\u003eEndTime: 结束日期 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBeginEndDate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e",
                        "schema": {
                            "$ref": "#/definitions/int"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/gettimestatement": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "注册汇总",
                "description": "一段时间内的注册人数按日期列表",
                "operationId": "/users/gettimestatement",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eBeginTime: 起始日期 \u003cbr/\u003eEndTime: 结束日期 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBeginEndDate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eRegisterNumber: \u003cbr/\u003eDateTime: ",
                        "schema": {
                            "$ref": "#/definitions/models.ResqUserRegisteInfo"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/gettimetotal": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "注册汇总",
                "description": "一段时间内的注册人数",
                "operationId": "/users/gettimetotal",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eBeginTime: 起始日期 \u003cbr/\u003eEndTime: 结束日期 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqBeginEndDate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/getuser": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "获取用户核心信息",
                "description": "获取用户核心信息",
                "operationId": "/users/getuser",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户信息详情\u003cbr/\u003eId: 用户id \u003cbr/\u003eState: 用户状态 \u003cbr/\u003ePlatform: 注册来源 \u003cbr/\u003ePwd: 密码 \u003cbr/\u003eSalt: 加密矢量 \u003cbr/\u003ePayPwd: 支付密码 \u003cbr/\u003ePaySalt: 支付密码加密矢量 \u003cbr/\u003eDeviceNumber: 设备驱动号 \u003cbr/\u003eVersion: 版本号 \u003cbr/\u003eAgent: useragent \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.Users"
                        }
                    },
                    "202": {
                        "description": "请求出错\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/getuseraccount": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "根据账号获取用户",
                "description": "根据账号获取用户",
                "operationId": "/users/getuseraccount",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"账号\"\u003cbr/\u003eAccount: 账号 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetUserAccount"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "已经注册\u003cbr/\u003eId: 账用户号id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eAccountType: 账号类型 \u003cbr/\u003eAccount: 账号 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.UserAccount"
                        }
                    },
                    "202": {
                        "description": "请求出错\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/getuserbyidcard": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "根据身份证号查询",
                "description": "根据身份证号查询",
                "operationId": "/users/getuserbyidcard",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"获取用户的银行卡\"\u003cbr/\u003eIdcard: 身份证 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserIdcard"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eUserId: 用户信息id \u003cbr/\u003eInviter: 邀请人 id \u003cbr/\u003eNickName: 用户名 \u003cbr/\u003eUserName: 真实姓名 \u003cbr/\u003eGender: 性别 \u003cbr/\u003eIdCard: 身份证号 \u003cbr/\u003ePic: 头像地址 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.UserInfos"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/getuserforsrarch": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "获取用户为搜索提供的",
                "description": "获取用户为搜索提供的",
                "operationId": "/users/getuserforsrarch",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"查询用户\"\u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqOffsetLimit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登录成功\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eUserName: 实名 \u003cbr/\u003eIdCard: 身份证号 \u003cbr/\u003eAccount: 账号 \u003cbr/\u003eCreateTime: 注册时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespUserForSrarch"
                        }
                    },
                    "202": {
                        "description": "登录失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/getuserinfo": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "查询用户详情",
                "description": "查询用户详情",
                "operationId": "/users/getuserinfo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户信息详情\u003cbr/\u003eUserId: 用户信息id \u003cbr/\u003eInviter: 邀请人 id \u003cbr/\u003eNickName: 用户名 \u003cbr/\u003eUserName: 真实姓名 \u003cbr/\u003eGender: 性别 \u003cbr/\u003eIdCard: 身份证号 \u003cbr/\u003ePic: 头像地址 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.UserInfos"
                        }
                    },
                    "202": {
                        "description": "请求出错\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/getuserinfobyuserid": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "获取用户相关信息",
                "description": "获取用户相关信息",
                "operationId": "/users/getuserinfobyuserid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户信息详情\u003cbr/\u003e",
                        "schema": {
                            "$ref": "#/definitions/models.RespUsers"
                        }
                    },
                    "202": {
                        "description": "请求出错\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/getuserinvited": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "获取用户邀请的人",
                "description": "获取用户邀请的人",
                "operationId": "/users/getuserinvited",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户邀请的人\u003cbr/\u003eUserId: 用户信息id \u003cbr/\u003eInviter: 邀请人 id \u003cbr/\u003eNickName: 用户名 \u003cbr/\u003eUserName: 真实姓名 \u003cbr/\u003eGender: 性别 \u003cbr/\u003eIdCard: 身份证号 \u003cbr/\u003ePic: 头像地址 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.UserInfos"
                        }
                    },
                    "202": {
                        "description": "请求出错\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/getuserinvitedinfo": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "获取用户邀请的人",
                "description": "获取用户邀请的人",
                "operationId": "/users/getuserinvitedinfo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"用户id\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户邀请的人\u003cbr/\u003eAccount: 账号 \u003cbr/\u003eUserId: 用户信息id \u003cbr/\u003eInviter: 邀请人 id \u003cbr/\u003eNickName: 用户名 \u003cbr/\u003eUserName: 真实姓名 \u003cbr/\u003eGender: 性别 \u003cbr/\u003eIdCard: 身份证号 \u003cbr/\u003ePic: 头像地址 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespUserInfoWithAccount"
                        }
                    },
                    "202": {
                        "description": "请求出错\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/getusermobile": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "获得用户的手机号",
                "description": "获得用户的手机号",
                "operationId": "/users/getusermobile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"查询用户的手机号\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登录成功\u003cbr/\u003eId: 账用户号id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eAccountType: 账号类型 \u003cbr/\u003eAccount: 账号 \u003cbr/\u003eCreateTime: 创建时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.UserAccount"
                        }
                    },
                    "202": {
                        "description": "登录失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/getuserssumbyquery": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "搜索用户",
                "description": "搜索用户",
                "operationId": "/users/getuserssumbyquery",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"查询用户\"\u003cbr/\u003eQuery: 查询 \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetUserByQuery"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登录成功\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eInviter: 邀请人id \u003cbr/\u003eNickName: 昵称 \u003cbr/\u003eUserName: 实名 \u003cbr/\u003eGender: 性别 \u003cbr/\u003eIdCard: 身份证号 \u003cbr/\u003eAccount: 账号 \u003cbr/\u003eAccountType: 账号类型 \u003cbr/\u003eAgent: 注册设备信息 \u003cbr/\u003ePlatform: 渠道 \u003cbr/\u003ePlatformName: 渠道名称 \u003cbr/\u003eCreateTime: 注册时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespUser"
                        }
                    },
                    "202": {
                        "description": "登录失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/getusersum": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "搜索用户",
                "description": "搜索用户",
                "operationId": "/users/getusersum",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"查询用户\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登录成功\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eInviter: 邀请人id \u003cbr/\u003eNickName: 昵称 \u003cbr/\u003eUserName: 实名 \u003cbr/\u003eGender: 性别 \u003cbr/\u003eIdCard: 身份证号 \u003cbr/\u003eAccount: 账号 \u003cbr/\u003eAccountType: 账号类型 \u003cbr/\u003eAgent: 注册设备信息 \u003cbr/\u003ePlatform: 渠道 \u003cbr/\u003ePlatformName: 渠道名称 \u003cbr/\u003eCreateTime: 注册时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespUser"
                        }
                    },
                    "202": {
                        "description": "登录失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/getusertag": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "获得用户的tag(key-value)",
                "description": "获得用户的tag(key-value)",
                "operationId": "/users/getusertag",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eKey: 用户标签键 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqGetUserTag"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e",
                        "schema": {
                            "$ref": "#/definitions/models.GetUserTagKV"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/getwithdrawbankcard": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "获取用户提现的银行卡",
                "description": "获取用户提现的银行卡",
                "operationId": "/users/getwithdrawbankcard",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"获取用户的银行卡\"\u003cbr/\u003eUserId: 用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserId"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回用户的余额和冻结余额\u003cbr/\u003eId: \u003cbr/\u003eUserId: \u003cbr/\u003eState: \u003cbr/\u003ePayWay: \u003cbr/\u003eTotalIn: \u003cbr/\u003eTotalOut: \u003cbr/\u003ePhone: \u003cbr/\u003eCard: \u003cbr/\u003eBankCode: \u003cbr/\u003eBindId: \u003cbr/\u003eBankName: \u003cbr/\u003eBankDesc: \u003cbr/\u003eCreateTime: \u003cbr/\u003eSingleLimit: \u003cbr/\u003eDayLimit: ",
                        "schema": {
                            "$ref": "#/definitions/models.RespBankcard"
                        }
                    },
                    "202": {
                        "description": "失败各种原因\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "用户登录",
                "description": "用户登录",
                "operationId": "/users/login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"账号\"\u003cbr/\u003eAccount: 账号 \u003cbr/\u003ePwd: 密码 \u003cbr/\u003eAgent: useragent \u003cbr/\u003eIp: ip ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登录成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "登录失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/registe": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "用户注册",
                "description": "用户注册",
                "operationId": "/users/registe",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"账号\"\u003cbr/\u003eAccount: 账号 \u003cbr/\u003ePwd: 密码 \u003cbr/\u003eDeviceNumber: 设备号 \u003cbr/\u003eVersion: 版本号 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqRegiste"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "创建成功并返回用户id\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "创建失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/setinviter": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "设置用户邀请人",
                "description": "设置用户邀请人",
                "operationId": "/users/setinviter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"设置用户邀请人id\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eInviter: 用户邀请人的用户id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqSetInviter"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "创建成功并返回用户id\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "创建失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/setpaypwd": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "设置用户交易密码",
                "description": "设置用户交易密码",
                "operationId": "/users/setpaypwd",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"交易密码\"\u003cbr/\u003e",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserPwd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "设置成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "设置失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/setpwd": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "设置用户登录密码",
                "description": "设置用户登录密码",
                "operationId": "/users/setpwd",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003e",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserPwd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "设置成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "设置失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/setuseragent": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "记录用户的来源",
                "description": "记录用户的来源",
                "operationId": "/users/setuseragent",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"获取用户的银行卡\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eAgent: useragent ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserAgent"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "记录用户的来源\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/setusername": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "设置用户身份证与姓名",
                "description": "设置用户身份证与姓名",
                "operationId": "/users/setusername",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"交易密码\"\u003cbr/\u003e",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserNameReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "设置成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "设置失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/setuserplatform": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "记录用户的来源",
                "description": "记录用户的来源",
                "operationId": "/users/setuserplatform",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"获取用户的银行卡\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003ePlatform: 平台id ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserPlatform"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "记录用户的来源\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/setuserstate": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "记录用户的来源",
                "description": "记录用户的来源",
                "operationId": "/users/setuserstate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"获取用户的银行卡\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eState: 状态 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserState"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "记录用户的来源\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/users/setusertag": {
            "post": {
                "tags": [
                    "/users"
                ],
                "summary": "设置用户的tag(key-value)",
                "description": "设置用户的tag(key-value)",
                "operationId": "/users/setusertag",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eKey: 用户标签键 \u003cbr/\u003eValue: 用户标签值 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqSetUserTag"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/withdrawrecord/apply": {
            "post": {
                "tags": [
                    "/withdrawrecord"
                ],
                "summary": "后台提现",
                "description": "后台提现",
                "operationId": "/withdrawrecord/apply",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eWithdrawId: 提现id \u003cbr/\u003eOrderNo: 订单号 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqWithdrawApply"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/withdrawrecord/getwithdrawbyid": {
            "post": {
                "tags": [
                    "/withdrawrecord"
                ],
                "summary": "提现记录查询",
                "description": "提现记录查询",
                "operationId": "/withdrawrecord/getwithdrawbyid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eWithdrawId: 提现id \u003cbr/\u003eOrderNo: 订单号 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqWithdrawApply"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eId: 提现记录id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eState: 状态 \u003cbr/\u003eBankcardId: 银行卡id \u003cbr/\u003eDrawNumber: 提现金额 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 \u003cbr/\u003eBankCode: 银行编号 \u003cbr/\u003eBankName: 银行名称 \u003cbr/\u003eUserName: 姓名 \u003cbr/\u003eAccount: 账号 \u003cbr/\u003eCard: 卡号 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespWithdrawRecord"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/withdrawrecord/getwithdrawrecordstatistics": {
            "post": {
                "tags": [
                    "/withdrawrecord"
                ],
                "summary": "提现记录",
                "description": "提现记录",
                "operationId": "/withdrawrecord/getwithdrawrecordstatistics",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eState: 状态 \u003cbr/\u003eBeginTime: 起始时间 \u003cbr/\u003eEndTime: 结束时间 \u003cbr/\u003eOffset: 偏移 \u003cbr/\u003eLimit: 条数 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqWithdrawUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eWithdrawId: 提现id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eAccount: 账号 \u003cbr/\u003eUserName: 实名 \u003cbr/\u003eDrawNumber: 提现金额 \u003cbr/\u003eBankCode: 银行编码 \u003cbr/\u003eBankName: 银行名称 \u003cbr/\u003eCard: 卡号 \u003cbr/\u003eCreateTime: 提现申请时间 \u003cbr/\u003eState: 提现状态 \u003cbr/\u003eUpdateTime: 状态更新时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.RespWithdrawRecordStatistics"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/withdrawrecord/processinglist": {
            "post": {
                "tags": [
                    "/withdrawrecord"
                ],
                "summary": "第三方处理中的提现列表",
                "description": "第三方处理中的提现列表",
                "operationId": "/withdrawrecord/processinglist",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eDuring: \u003cbr/\u003eNumber: ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqProcDrawRecord"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e",
                        "schema": {
                            "$ref": "#/definitions/mdoels.RespWithdrawRecordList"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/withdrawrecord/record": {
            "post": {
                "tags": [
                    "/withdrawrecord"
                ],
                "summary": "提现记录",
                "description": "提现记录",
                "operationId": "/withdrawrecord/record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eOffset: 索引偏移 \u003cbr/\u003eLimit: 获取数量 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqWithdrawRecord"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "提现记录\u003cbr/\u003eId: 提现记录id \u003cbr/\u003eUserId: 用户id \u003cbr/\u003eState: 提现状态 \u003cbr/\u003eBankcardId: 银行卡id \u003cbr/\u003eDrawNumber: 提现金额 \u003cbr/\u003eCreateTime: 创建时间 \u003cbr/\u003eUpdateTime: 更新时间 ",
                        "schema": {
                            "$ref": "#/definitions/models.WithdrawRecord"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/withdrawrecord/withdrawbyid": {
            "post": {
                "tags": [
                    "/withdrawrecord"
                ],
                "summary": "提现",
                "description": "提现",
                "operationId": "/withdrawrecord/withdrawbyid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eWithdrawId: 提现id \u003cbr/\u003eOrderNo: 订单号 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqWithdrawApply"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/withdrawrecord/withdrawcount": {
            "post": {
                "tags": [
                    "/withdrawrecord"
                ],
                "summary": "提现",
                "description": "提现",
                "operationId": "/withdrawrecord/withdrawcount",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eUserId: 用户id \u003cbr/\u003eWithdrawTime: 提现时间 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqWithdrawCount"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/withdrawrecord/withdrawlist": {
            "post": {
                "tags": [
                    "/withdrawrecord"
                ],
                "summary": "提现",
                "description": "提现",
                "operationId": "/withdrawrecord/withdrawlist",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eTimeBegin: \u003cbr/\u003eTimeEnd: \u003cbr/\u003eUserId: \u003cbr/\u003eNumberBegin: \u003cbr/\u003eNumberEnd: \u003cbr/\u003eDrawState: \u003cbr/\u003eOffset: \u003cbr/\u003eLimit: ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqDrawRecord"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eId: \u003cbr/\u003eUserId: \u003cbr/\u003eAccount: \u003cbr/\u003eUserType: \u003cbr/\u003eDrawNumber: \u003cbr/\u003eCheckState: \u003cbr/\u003eState: \u003cbr/\u003eCard: \u003cbr/\u003eCreateTime: \u003cbr/\u003eCheckTime: ",
                        "schema": {
                            "$ref": "#/definitions/models.RespDrawRecord"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                },
                "deprecated": true
            }
        },
        "/withdrawrecord/withdrawlistcount": {
            "post": {
                "tags": [
                    "/withdrawrecord"
                ],
                "summary": "提现",
                "description": "提现",
                "operationId": "/withdrawrecord/withdrawlistcount",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eTimeBegin: \u003cbr/\u003eTimeEnd: \u003cbr/\u003eUserId: \u003cbr/\u003eNumberBegin: \u003cbr/\u003eNumberEnd: \u003cbr/\u003eDrawState: \u003cbr/\u003eOffset: \u003cbr/\u003eLimit: ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqDrawRecord"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003eTimeBegin: \u003cbr/\u003eTimeEnd: \u003cbr/\u003eUserId: \u003cbr/\u003eNumberBegin: \u003cbr/\u003eNumberEnd: \u003cbr/\u003eDrawState: \u003cbr/\u003eOffset: \u003cbr/\u003eLimit: ",
                        "schema": {
                            "$ref": "#/definitions/models.ReqDrawRecord"
                        }
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/withdrawrecord/withdrawrecordupdate": {
            "post": {
                "tags": [
                    "/withdrawrecord"
                ],
                "summary": "提现处理",
                "description": "提现处理",
                "operationId": "/withdrawrecord/withdrawrecordupdate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eWithdrawId: 提现id \u003cbr/\u003eMark: 标记 \u003cbr/\u003eState: 状态 ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqWithdrawState"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        },
        "/withdrawrecord/withdrawstateupdate": {
            "post": {
                "tags": [
                    "/withdrawrecord"
                ],
                "summary": "提现批量处理",
                "description": "提现批量处理",
                "operationId": "/withdrawrecord/withdrawstateupdate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "\"\"\u003cbr/\u003eState: id/状态 map[string]int json ",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqDrawState"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功\u003cbr/\u003e"
                    },
                    "202": {
                        "description": "请求失败\u003cbr/\u003e"
                    }
                }
            }
        }
    },
    "definitions": {
        "": {
            "type": "object"
        },
        "apim.Message": {
            "title": "apim.Message",
            "type": "object"
        },
        "int": {
            "title": "int",
            "type": "object"
        },
        "map[string]string": {
            "title": "map[string]string",
            "type": "object"
        },
        "mdoels.RespWithdrawRecordList": {
            "title": "mdoels.RespWithdrawRecordList",
            "type": "object"
        },
        "models.Balance": {
            "title": "models.Balance",
            "type": "object",
            "properties": {
                "FreezeNumber": {
                    "description": "冻结收益 ",
                    "type": "integer",
                    "format": "int64"
                },
                "HistoryDemandNumber": {
                    "description": "活期历史收益 ",
                    "type": "integer",
                    "format": "int64"
                },
                "HistoryNormalNumber": {
                    "description": "定期历史收益 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Id": {
                    "description": "用户余额id ",
                    "type": "integer",
                    "format": "int64"
                },
                "Number": {
                    "description": "余额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UpdateTime": {
                    "description": "更新时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.Bankcard": {
            "title": "models.Bankcard",
            "type": "object",
            "properties": {
                "BankCode": {
                    "description": "银行代号 ",
                    "type": "string"
                },
                "BankDesc": {
                    "description": "银行描述 ",
                    "type": "string"
                },
                "BankName": {
                    "description": "银行名 ",
                    "type": "string"
                },
                "BindId": {
                    "description": "绑定id ",
                    "type": "string"
                },
                "Card": {
                    "description": "银行卡号 ",
                    "type": "string"
                },
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Id": {
                    "description": "银行卡id ",
                    "type": "integer",
                    "format": "int64"
                },
                "PayWay": {
                    "description": "支付方式 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Phone": {
                    "description": "预留手机号 ",
                    "type": "string"
                },
                "State": {
                    "description": "状态 ",
                    "type": "integer",
                    "format": "int64"
                },
                "TotalIn": {
                    "description": "最大流入限额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "TotalOut": {
                    "description": "最大流出限额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.Banner": {
            "title": "models.Banner",
            "type": "object",
            "properties": {
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Id": {
                    "description": "轮播id ",
                    "type": "integer",
                    "format": "int64"
                },
                "Pic": {
                    "description": "图片地址 ",
                    "type": "string"
                },
                "Platform": {
                    "description": "平台id ",
                    "type": "integer",
                    "format": "int64"
                },
                "State": {
                    "description": "状态 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Url": {
                    "description": "点击跳转地址 ",
                    "type": "string"
                }
            }
        },
        "models.Broadcast": {
            "title": "models.Broadcast",
            "type": "object",
            "properties": {
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Id": {
                    "description": "广播id ",
                    "type": "integer",
                    "format": "int64"
                },
                "Message": {
                    "description": "信息 ",
                    "type": "string"
                },
                "State": {
                    "description": "状态 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.CardLimit": {
            "title": "models.CardLimit",
            "type": "object",
            "properties": {
                "BankCode": {
                    "description": "银行代号 ",
                    "type": "string"
                },
                "BankName": {
                    "description": "银行名字 ",
                    "type": "string"
                },
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "DayLimit": {
                    "description": "天数 ",
                    "type": "string"
                },
                "Id": {
                    "description": "银行卡限制id ",
                    "type": "integer",
                    "format": "int64"
                },
                "PayWay": {
                    "description": "支付方式 ",
                    "type": "integer",
                    "format": "int64"
                },
                "SingleLimit": {
                    "description": "单笔限额 ",
                    "type": "string"
                }
            }
        },
        "models.ChargeRecord": {
            "title": "models.ChargeRecord",
            "type": "object",
            "properties": {
                "BankcardId": {
                    "description": "银行卡id ",
                    "type": "integer",
                    "format": "int64"
                },
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Id": {
                    "description": "充值记录id ",
                    "type": "integer",
                    "format": "int64"
                },
                "Number": {
                    "description": "支付金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "PayNo": {
                    "description": "订单号 ",
                    "type": "string"
                },
                "PayWay": {
                    "description": "支付方式 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.Config": {
            "title": "models.Config",
            "type": "object",
            "properties": {
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Desc": {
                    "description": "描述 ",
                    "type": "string"
                },
                "Id": {
                    "description": "配置id ",
                    "type": "integer",
                    "format": "int64"
                },
                "Key": {
                    "description": "键 ",
                    "type": "string"
                },
                "UpdateTime": {
                    "description": "更新时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Value": {
                    "description": "值 ",
                    "type": "string"
                }
            }
        },
        "models.Coupon": {
            "title": "models.Coupon",
            "type": "object",
            "properties": {
                "CouponDesc": {
                    "description": "优惠券信息 ",
                    "type": "string"
                },
                "CouponName": {
                    "description": "优惠券描述 ",
                    "type": "string"
                },
                "CouponRule": {
                    "description": "优惠券规则 ",
                    "type": "string"
                },
                "CouponType": {
                    "description": "优惠券类型  0 | CouponTypeDiscount | CouponTypeRaise ",
                    "type": "integer",
                    "format": "int64"
                },
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Id": {
                    "description": "优惠券类型id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.DividendSubs": {
            "title": "models.DividendSubs",
            "type": "object"
        },
        "models.Experience": {
            "title": "models.Experience",
            "type": "object",
            "properties": {
                "CheckDate": {
                    "description": "回款日 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EarningsNumber": {
                    "description": "回款金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Id": {
                    "description": "体验金id ",
                    "type": "integer",
                    "format": "int64"
                },
                "Message": {
                    "description": "描述 ",
                    "type": "string"
                },
                "Number": {
                    "description": "金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Status": {
                    "description": "状态 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.GetUserTagKV": {
            "title": "models.GetUserTagKV",
            "type": "object"
        },
        "models.HistoryRecords": {
            "title": "models.HistoryRecords",
            "type": "object"
        },
        "models.Message": {
            "title": "models.Message",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Id": {
                    "description": "信息id ",
                    "type": "integer",
                    "format": "int64"
                },
                "Msg": {
                    "description": "信息 ",
                    "type": "string"
                },
                "State": {
                    "description": "信息状态 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Title": {
                    "description": "消息标题 ",
                    "type": "string"
                },
                "Url": {
                    "description": "点击后跳转链接 ",
                    "type": "string"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.Platform": {
            "title": "models.Platform",
            "type": "object",
            "properties": {
                "CreateTime": {
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Id": {
                    "type": "integer",
                    "format": "int64"
                },
                "Name": {
                    "type": "string"
                },
                "PackCode": {
                    "type": "string"
                },
                "PlatformId": {
                    "type": "integer",
                    "format": "int64"
                },
                "State": {
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.PurchaseRecord": {
            "title": "models.PurchaseRecord",
            "type": "object",
            "properties": {
                "BidId": {
                    "description": "标的id ",
                    "type": "integer",
                    "format": "int64"
                },
                "CouponRules": {
                    "description": "购买时使用的券 ",
                    "type": "string"
                },
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Id": {
                    "type": "integer",
                    "format": "int64"
                },
                "Number": {
                    "description": "购买金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "PayNumber": {
                    "description": "支付金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "PurchaseRecordPoolId": {
                    "description": "购买记录池id ",
                    "type": "integer",
                    "format": "int64"
                },
                "RemainNumber": {
                    "description": "剩余金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UpdateTime": {
                    "description": "更新时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.PurchaseRecordPool": {
            "title": "models.PurchaseRecordPool",
            "type": "object",
            "properties": {
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Id": {
                    "description": "购买池id ",
                    "type": "integer",
                    "format": "int64"
                },
                "Number": {
                    "description": "购买金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "PayNumber": {
                    "description": "支付金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RemainNumber": {
                    "description": "持有金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Status": {
                    "description": "持有状态 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UpdateTime": {
                    "description": "更新时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqAddBalance": {
            "title": "models.ReqAddBalance",
            "type": "object",
            "properties": {
                "BankcardId": {
                    "description": "银行卡号 ",
                    "type": "integer",
                    "format": "int64"
                },
                "FromId": {
                    "description": "来自id ",
                    "type": "integer",
                    "format": "int64"
                },
                "FromType": {
                    "description": "来自类型 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Number": {
                    "description": "金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "PayNo": {
                    "description": "支付订单号 ",
                    "type": "string"
                },
                "PayWay": {
                    "description": "支付方式类型 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RecordDesc": {
                    "description": "来自描述 ",
                    "type": "string"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqAddBid": {
            "title": "models.ReqAddBid",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "开卖时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "BidDesc": {
                    "description": "标描述 ",
                    "type": "string"
                },
                "BidInfoPath": {
                    "description": "标的详情地址 ",
                    "type": "string"
                },
                "BidMark": {
                    "description": "标的购买限制 ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidName": {
                    "description": "标名称 ",
                    "type": "string"
                },
                "BidNo": {
                    "description": "标订单号 ",
                    "type": "string"
                },
                "BidRate": {
                    "description": "标的利率 ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidRateShow": {
                    "description": "标利率特殊显示 ",
                    "type": "string"
                },
                "CheckDate": {
                    "description": "结算日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "DaysLimit": {
                    "description": "这是多少天标 ",
                    "type": "integer",
                    "format": "int64"
                },
                "EndTime": {
                    "description": "结束时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "MaxBuySize": {
                    "description": "最大购买笔数 ",
                    "type": "integer",
                    "format": "int64"
                },
                "MaxNumber": {
                    "description": "最大金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "MinNumber": {
                    "description": "起投金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillBidMark": {
                    "description": "标的购买限制 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillBidType": {
                    "description": "标计息类型 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillDaysLimit": {
                    "description": "标的物最大期限 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillRemainNumberMin": {
                    "description": "全部符合条件的标 累计小于最小续标金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "StepNumber": {
                    "description": "增幅金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "TotalNumber": {
                    "description": "标总金额 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqAddBidNew": {
            "title": "models.ReqAddBidNew",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "开卖时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "BidDesc": {
                    "description": "标描述 ",
                    "type": "string"
                },
                "BidInfoPath": {
                    "description": "标的详情地址 ",
                    "type": "string"
                },
                "BidMark": {
                    "description": "标的购买限制 ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidName": {
                    "description": "标名称 ",
                    "type": "string"
                },
                "BidNo": {
                    "description": "标订单号 ",
                    "type": "string"
                },
                "BidRate": {
                    "description": "标的利率 ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidRateShow": {
                    "description": "标利率特殊显示 ",
                    "type": "string"
                },
                "BidType": {
                    "description": "标的类型 ",
                    "type": "integer",
                    "format": "int64"
                },
                "DaysLimit": {
                    "description": "这是多少天标 ",
                    "type": "integer",
                    "format": "int64"
                },
                "EndTime": {
                    "description": "结束时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "MaxBuySize": {
                    "description": "最大购买笔数 ",
                    "type": "integer",
                    "format": "int64"
                },
                "MaxNumber": {
                    "description": "最大金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "MinNumber": {
                    "description": "起投金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillBidMark": {
                    "description": "标的购买限制 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillBidType": {
                    "description": "标计息类型 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillDaysLimit": {
                    "description": "标的物最大期限 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillRemainNumberMin": {
                    "description": "全部符合条件的标 累计小于最小续标金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "StepNumber": {
                    "description": "增幅金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "TotalNumber": {
                    "description": "标总金额 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqAddBidPool": {
            "title": "models.ReqAddBidPool",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "开卖时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "BidDesc": {
                    "description": "标描述 ",
                    "type": "string"
                },
                "BidInfoPath": {
                    "description": "标的详情地址 ",
                    "type": "string"
                },
                "BidName": {
                    "description": "标名称 ",
                    "type": "string"
                },
                "BidNo": {
                    "description": "标订单号 ",
                    "type": "string"
                },
                "CheckDate": {
                    "description": "结算日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "DaysLimit": {
                    "description": "这是多少天标 ",
                    "type": "integer",
                    "format": "int64"
                },
                "EndTime": {
                    "description": "结束时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "TotalNumber": {
                    "description": "标总金额 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqAddBidSatisfy": {
            "title": "models.ReqAddBidSatisfy",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "开卖时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "BidDesc": {
                    "description": "标描述 ",
                    "type": "string"
                },
                "BidInfoPath": {
                    "description": "标的详情地址 ",
                    "type": "string"
                },
                "BidMark": {
                    "description": "标的购买限制 ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidName": {
                    "description": "标名称 ",
                    "type": "string"
                },
                "BidNo": {
                    "description": "标订单号 ",
                    "type": "string"
                },
                "BidRate": {
                    "description": "标的利率 ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidRateShow": {
                    "description": "标利率特殊显示 ",
                    "type": "string"
                },
                "DaysLimit": {
                    "description": "这是多少天标 ",
                    "type": "integer",
                    "format": "int64"
                },
                "EndTime": {
                    "description": "结束时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "MaxBuySize": {
                    "description": "最大购买笔数 ",
                    "type": "integer",
                    "format": "int64"
                },
                "MaxNumber": {
                    "description": "最大金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "MinNumber": {
                    "description": "起投金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillBidMark": {
                    "description": "标的购买限制 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillBidType": {
                    "description": "标计息类型 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillDaysLimit": {
                    "description": "标的物最大期限 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillRemainNumberMin": {
                    "description": "全部符合条件的标 累计小于最小续标金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "StepNumber": {
                    "description": "增幅金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "TotalNumber": {
                    "description": "标总金额 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqAddCard": {
            "title": "models.ReqAddCard",
            "type": "object",
            "properties": {
                "BankCode": {
                    "description": "银行编码 ",
                    "type": "string"
                },
                "BankName": {
                    "description": "银行名称 ",
                    "type": "string"
                },
                "CardNumber": {
                    "type": "string"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqAddDividendSub": {
            "title": "models.ReqAddDividendSub",
            "type": "object",
            "properties": {
                "BeginDate": {
                    "description": "开始日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "BidId": {
                    "description": "标的id ",
                    "type": "integer",
                    "format": "int64"
                },
                "Rate": {
                    "description": "利率 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqAddExperience": {
            "title": "models.ReqAddExperience",
            "type": "object",
            "properties": {
                "Message": {
                    "description": "描述 ",
                    "type": "string"
                },
                "Number": {
                    "description": "金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqAddPlatform": {
            "title": "models.ReqAddPlatform",
            "type": "object",
            "properties": {
                "Name": {
                    "description": "渠道名称 ",
                    "type": "string"
                },
                "PackCode": {
                    "description": "包名 ",
                    "type": "string"
                },
                "PlatformId": {
                    "description": "渠道id ",
                    "type": "integer",
                    "format": "int64"
                },
                "State": {
                    "description": "状态 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqAddPurchase": {
            "title": "models.ReqAddPurchase",
            "type": "object",
            "properties": {
                "BidId": {
                    "description": "标的id ",
                    "type": "integer",
                    "format": "int64"
                },
                "CouponRules": {
                    "description": "购买时使用的券 ",
                    "type": "string"
                },
                "Number": {
                    "description": "购买金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqAddPurchasePool": {
            "title": "models.ReqAddPurchasePool",
            "type": "object",
            "properties": {
                "Number": {
                    "description": "购买金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqAddReturned": {
            "title": "models.ReqAddReturned",
            "type": "object",
            "properties": {
                "PurchaseRecordId": {
                    "description": "购买记录id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqAddReturnedPool": {
            "title": "models.ReqAddReturnedPool",
            "type": "object",
            "properties": {
                "PurchaseRecordPoolId": {
                    "description": "购买记录id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqAddUserCoupon": {
            "title": "models.ReqAddUserCoupon",
            "type": "object",
            "properties": {
                "CouponId": {
                    "description": "优惠券类型 id ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqAllMessageAdd": {
            "title": "models.ReqAllMessageAdd",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "发布时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Msg": {
                    "description": "消息 ",
                    "type": "string"
                },
                "Title": {
                    "description": "标题 ",
                    "type": "string"
                },
                "Url": {
                    "description": "跳转链接 ",
                    "type": "string"
                }
            }
        },
        "models.ReqBankcard": {
            "title": "models.ReqBankcard",
            "type": "object",
            "properties": {
                "BankCode": {
                    "description": "银行编码 ",
                    "type": "string"
                },
                "BankName": {
                    "description": "银行名称 ",
                    "type": "string"
                },
                "BindId": {
                    "description": "绑定id ",
                    "type": "string"
                },
                "Card": {
                    "description": "卡号 ",
                    "type": "string"
                },
                "PayWay": {
                    "description": "支付渠道 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Phone": {
                    "description": "预留手机号 ",
                    "type": "string"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqBankcardCard": {
            "title": "models.ReqBankcardCard",
            "type": "object",
            "properties": {
                "Card": {
                    "description": "银行卡号 ",
                    "type": "string"
                }
            }
        },
        "models.ReqBankcardId": {
            "title": "models.ReqBankcardId",
            "type": "object",
            "properties": {
                "BankcardId": {
                    "description": "银行卡id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqBankcardState": {
            "title": "models.ReqBankcardState",
            "type": "object",
            "properties": {
                "BankcardId": {
                    "description": "银行卡id ",
                    "type": "integer",
                    "format": "int64"
                },
                "State": {
                    "description": "银行卡号 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqBannerAdd": {
            "title": "models.ReqBannerAdd",
            "type": "object",
            "properties": {
                "Pic": {
                    "description": "图片url ",
                    "type": "string"
                },
                "Platform": {
                    "description": "平台id ",
                    "type": "integer",
                    "format": "int64"
                },
                "State": {
                    "description": "轮播图状态号 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Url": {
                    "description": "跳转url ",
                    "type": "string"
                }
            }
        },
        "models.ReqBannerList": {
            "title": "models.ReqBannerList",
            "type": "object",
            "properties": {
                "Limit": {
                    "description": "获取数量 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "索引偏移 ",
                    "type": "integer",
                    "format": "int64"
                },
                "State": {
                    "description": "状态 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqBannerUpdate": {
            "title": "models.ReqBannerUpdate",
            "type": "object",
            "properties": {
                "Id": {
                    "description": "轮播图id ",
                    "type": "integer",
                    "format": "int64"
                },
                "Pic": {
                    "description": "图片url ",
                    "type": "string"
                },
                "Platform": {
                    "description": "平台id ",
                    "type": "integer",
                    "format": "int64"
                },
                "State": {
                    "description": "轮播图状态号 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Url": {
                    "description": "跳转url ",
                    "type": "string"
                }
            }
        },
        "models.ReqBeginEndDate": {
            "title": "models.ReqBeginEndDate",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "起始日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EndTime": {
                    "description": "结束日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                }
            }
        },
        "models.ReqBidId": {
            "title": "models.ReqBidId",
            "type": "object",
            "properties": {
                "BidId": {
                    "description": "标的id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqCardInfo": {
            "title": "models.ReqCardInfo",
            "type": "object",
            "properties": {
                "BankCode": {
                    "description": "银行编码 ",
                    "type": "string"
                },
                "PayWay": {
                    "description": "渠道 ",
                    "type": "string"
                }
            }
        },
        "models.ReqChangeBidType": {
            "title": "models.ReqChangeBidType",
            "type": "object",
            "properties": {
                "BidId": {
                    "description": "标的id ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidType": {
                    "description": "标的类型 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqChargeRecord": {
            "title": "models.ReqChargeRecord",
            "type": "object",
            "properties": {
                "Bankcard": {
                    "description": "银行卡号 ",
                    "type": "string"
                },
                "Mobile": {
                    "description": "手机号 ",
                    "type": "string"
                },
                "Number": {
                    "description": "金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqCouponId": {
            "title": "models.ReqCouponId",
            "type": "object",
            "properties": {
                "CouponId": {
                    "description": "优惠券类型 id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqDate": {
            "title": "models.ReqDate",
            "type": "object",
            "properties": {
                "Date": {
                    "description": "日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                }
            }
        },
        "models.ReqDateOffsetLimit": {
            "title": "models.ReqDateOffsetLimit",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "起始日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EndTime": {
                    "description": "结束日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Limit": {
                    "description": "限制取多少条 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "起始偏移 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqDeviceNumber": {
            "title": "models.ReqDeviceNumber",
            "type": "object",
            "properties": {
                "DeviceNumber": {
                    "description": "设备号 ",
                    "type": "string"
                }
            }
        },
        "models.ReqDrawRecord": {
            "title": "models.ReqDrawRecord",
            "type": "object",
            "properties": {
                "DrawState": {
                    "type": "integer",
                    "format": "int64"
                },
                "Limit": {
                    "type": "integer",
                    "format": "int64"
                },
                "NumberBegin": {
                    "type": "integer",
                    "format": "int64"
                },
                "NumberEnd": {
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "type": "integer",
                    "format": "int64"
                },
                "TimeBegin": {
                    "type": "string"
                },
                "TimeEnd": {
                    "type": "string"
                },
                "UserId": {
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqDrawState": {
            "title": "models.ReqDrawState",
            "type": "object",
            "properties": {
                "State": {
                    "description": "id/状态 map[string]int json ",
                    "type": "string"
                }
            }
        },
        "models.ReqFeedback": {
            "title": "models.ReqFeedback",
            "type": "object",
            "properties": {
                "Content": {
                    "description": "内容 ",
                    "type": "string"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqFullBids": {
            "title": "models.ReqFullBids",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "发布时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "BidName": {
                    "description": "标名 ",
                    "type": "string"
                },
                "CheckBeginTime": {
                    "description": "回款时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "CheckEndTime": {
                    "description": "汇款时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EndTime": {
                    "description": "发布时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "IncomeBeginTime": {
                    "description": "起息时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "IncomeEndTime": {
                    "description": "起息时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                }
            }
        },
        "models.ReqGetBalanceRecord": {
            "title": "models.ReqGetBalanceRecord",
            "type": "object",
            "properties": {
                "FromType": {
                    "description": "类型 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Limit": {
                    "description": "获取数量 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "索引偏移 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqGetBidsByQuery": {
            "title": "models.ReqGetBidsByQuery",
            "type": "object",
            "properties": {
                "BidMark": {
                    "description": "标的分类 ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidType": {
                    "description": "标的类型 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Limit": {
                    "description": "获取数量 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "索引偏移 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Query": {
                    "description": "搜索 ",
                    "type": "string"
                }
            }
        },
        "models.ReqGetBidsByQueryCountNew": {
            "title": "models.ReqGetBidsByQueryCountNew",
            "type": "object",
            "properties": {
                "BidMark": {
                    "description": "标的分类 ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidType": {
                    "description": "标的类型 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Query": {
                    "description": "搜索 ",
                    "type": "string"
                }
            }
        },
        "models.ReqGetBidsByQueryNew": {
            "title": "models.ReqGetBidsByQueryNew",
            "type": "object",
            "properties": {
                "BidMark": {
                    "description": "标的分类 ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidType": {
                    "description": "标的类型 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Limit": {
                    "description": "获取数量 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "索引偏移 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Query": {
                    "description": "搜索 ",
                    "type": "string"
                }
            }
        },
        "models.ReqGetBidsNew": {
            "title": "models.ReqGetBidsNew",
            "type": "object"
        },
        "models.ReqGetChargeRecordByPayNo": {
            "title": "models.ReqGetChargeRecordByPayNo",
            "type": "object",
            "properties": {
                "PayNo": {
                    "description": "支付订单号 ",
                    "type": "string"
                }
            }
        },
        "models.ReqGetDividendSub": {
            "title": "models.ReqGetDividendSub",
            "type": "object",
            "properties": {
                "BidId": {
                    "description": "标的id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqGetEstCheckReturned": {
            "title": "models.ReqGetEstCheckReturned",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "起始日期 可选 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EndTime": {
                    "description": "结束日期 可选 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "PurchaseRecordId": {
                    "description": "购买记录 id ",
                    "type": "integer",
                    "format": "int64"
                },
                "Real": {
                    "description": "真实数据 ",
                    "type": "boolean"
                },
                "Status": {
                    "description": "是否已经回款 可选 0 | ReturnedStateBefore | ReturnedStateAfter ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqGetMessages": {
            "title": "models.ReqGetMessages",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "开始时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EndTime": {
                    "description": "结束时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Limit": {
                    "description": "获取数量 ",
                    "type": "integer",
                    "format": "int64"
                },
                "MessageId": {
                    "description": "id ",
                    "type": "integer",
                    "format": "int64"
                },
                "MsgType": {
                    "description": "消息类型 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "索引偏移 ",
                    "type": "integer",
                    "format": "int64"
                },
                "State": {
                    "description": "状态 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Title": {
                    "description": "标题 ",
                    "type": "string"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqGetPurchaseFullCount": {
            "title": "models.ReqGetPurchaseFullCount",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "开始时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EndTime": {
                    "description": "结束时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Number": {
                    "description": "提现金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqGetPurchaseInfoTotal": {
            "title": "models.ReqGetPurchaseInfoTotal",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                },
                "BeginTime": {
                    "description": "起始日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "BidName": {
                    "description": "标名 ",
                    "type": "string"
                },
                "BidNo": {
                    "description": "标的编号 ",
                    "type": "string"
                },
                "DaysLimitMax": {
                    "description": "天数最大 ",
                    "type": "integer",
                    "format": "int64"
                },
                "DaysLimitMin": {
                    "description": "天数最小 ",
                    "type": "integer",
                    "format": "int64"
                },
                "EndTime": {
                    "description": "结束日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Platform": {
                    "description": "渠道id ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserName": {
                    "description": "用户真实姓名 ",
                    "type": "string"
                }
            }
        },
        "models.ReqGetPurchaseRecordByUserId": {
            "title": "models.ReqGetPurchaseRecordByUserId",
            "type": "object",
            "properties": {
                "Limit": {
                    "description": "获取数量 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "起始偏移 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqGetRanking": {
            "title": "models.ReqGetRanking",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "开始时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EndTime": {
                    "description": "结束时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Limit": {
                    "description": "获取数量 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "索引偏移 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqGetRankingByDay": {
            "title": "models.ReqGetRankingByDay",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "开始时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EndTime": {
                    "description": "结束时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Limit": {
                    "description": "获取数量 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "索引偏移 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Query": {
                    "description": "搜索 ",
                    "type": "string"
                }
            }
        },
        "models.ReqGetRecordHistory": {
            "title": "models.ReqGetRecordHistory",
            "type": "object",
            "properties": {
                "Limit": {
                    "description": "获取数量 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "索引偏移 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqGetReturned": {
            "title": "models.ReqGetReturned",
            "type": "object",
            "properties": {
                "BeginDate": {
                    "description": "开始日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "BidId": {
                    "description": "标的id ",
                    "type": "integer",
                    "format": "int64"
                },
                "EndDate": {
                    "description": "结束日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "PurchaseRecordId": {
                    "description": "购买记录id ",
                    "type": "integer",
                    "format": "int64"
                },
                "State": {
                    "description": "状态 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqGetUserAccount": {
            "title": "models.ReqGetUserAccount",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                }
            }
        },
        "models.ReqGetUserByQuery": {
            "title": "models.ReqGetUserByQuery",
            "type": "object",
            "properties": {
                "Limit": {
                    "description": "获取数量 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "索引偏移 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Query": {
                    "description": "查询 ",
                    "type": "string"
                }
            }
        },
        "models.ReqGetUserCoupon": {
            "title": "models.ReqGetUserCoupon",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "起始时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EndTime": {
                    "description": "截止时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqGetUserCouponCanUse": {
            "title": "models.ReqGetUserCouponCanUse",
            "type": "object",
            "properties": {
                "BidId": {
                    "description": "标的id ",
                    "type": "integer",
                    "format": "int64"
                },
                "Number": {
                    "description": "金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqGetUserCouponNew": {
            "title": "models.ReqGetUserCouponNew",
            "type": "object",
            "properties": {
                "CouponId": {
                    "description": "优惠券id 可选 ",
                    "type": "integer",
                    "format": "int64"
                },
                "CouponType": {
                    "description": "优惠券类型 0 | CouponTypeDiscount | CouponTypeRaise ",
                    "type": "integer",
                    "format": "int64"
                },
                "Limit": {
                    "description": "获取数量 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "索引偏移 ",
                    "type": "integer",
                    "format": "int64"
                },
                "State": {
                    "description": "优惠券状态 可选 0 | UserCouponStateUnused | UserCouponStateUsed | UserCouponStateOverdue ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id 可选 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqGetUserTag": {
            "title": "models.ReqGetUserTag",
            "type": "object",
            "properties": {
                "Key": {
                    "description": "用户标签键 ",
                    "type": "string"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqLogin": {
            "title": "models.ReqLogin",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                },
                "Agent": {
                    "description": "useragent ",
                    "type": "string"
                },
                "Ip": {
                    "description": "ip ",
                    "type": "string"
                },
                "Pwd": {
                    "description": "密码 ",
                    "type": "string"
                }
            }
        },
        "models.ReqMessageAdd": {
            "title": "models.ReqMessageAdd",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "开始时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Msg": {
                    "description": "消息 ",
                    "type": "string"
                },
                "Title": {
                    "description": "标题 ",
                    "type": "string"
                },
                "Url": {
                    "description": "跳转链接 ",
                    "type": "string"
                }
            }
        },
        "models.ReqMessageAddByAccount": {
            "title": "models.ReqMessageAddByAccount",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "发布时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Msg": {
                    "description": "消息 ",
                    "type": "string"
                },
                "Title": {
                    "description": "标题 ",
                    "type": "string"
                },
                "Url": {
                    "description": "跳转链接 ",
                    "type": "string"
                }
            }
        },
        "models.ReqMessageCount": {
            "title": "models.ReqMessageCount",
            "type": "object",
            "properties": {
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqMessageInfo": {
            "title": "models.ReqMessageInfo",
            "type": "object",
            "properties": {
                "MessageId": {
                    "description": "消息id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqMessageList": {
            "title": "models.ReqMessageList",
            "type": "object",
            "properties": {
                "Limit": {
                    "description": "获取数量 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "索引偏移 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqMessageRead": {
            "title": "models.ReqMessageRead",
            "type": "object",
            "properties": {
                "Id": {
                    "description": "消息id ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqMessageUpdate": {
            "title": "models.ReqMessageUpdate",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "发布时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "MessageId": {
                    "description": "消息id ",
                    "type": "integer",
                    "format": "int64"
                },
                "Msg": {
                    "description": "消息 ",
                    "type": "string"
                },
                "Title": {
                    "description": "标题 ",
                    "type": "string"
                },
                "Url": {
                    "description": "跳转链接 ",
                    "type": "string"
                }
            }
        },
        "models.ReqOffsetLimit": {
            "title": "models.ReqOffsetLimit",
            "type": "object",
            "properties": {
                "Limit": {
                    "description": "获取数量 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "索引偏移 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqPlatform": {
            "title": "models.ReqPlatform",
            "type": "object",
            "properties": {
                "Name": {
                    "description": "渠道名称 ",
                    "type": "string"
                },
                "PackCode": {
                    "description": "包名 ",
                    "type": "string"
                },
                "PlatformId": {
                    "description": "渠道id ",
                    "type": "integer",
                    "format": "int64"
                },
                "State": {
                    "description": "状态 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqPlatformPack": {
            "title": "models.ReqPlatformPack",
            "type": "object",
            "properties": {
                "PackCode": {
                    "description": "包名 ",
                    "type": "string"
                }
            }
        },
        "models.ReqPlatformPurchase": {
            "title": "models.ReqPlatformPurchase",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                },
                "BeginTime": {
                    "description": "投资起始日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EndTime": {
                    "description": "投资结束日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Limit": {
                    "description": "限制取多少条 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "起始偏移 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserName": {
                    "description": "姓名 ",
                    "type": "string"
                }
            }
        },
        "models.ReqPlatformState": {
            "title": "models.ReqPlatformState",
            "type": "object",
            "properties": {
                "Id": {
                    "description": "渠道设置id ",
                    "type": "integer",
                    "format": "int64"
                },
                "State": {
                    "description": "状态 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqPlatformStatistics": {
            "title": "models.ReqPlatformStatistics",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EndTime": {
                    "description": "时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "PlatformId": {
                    "description": "渠道id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqPlatformUserTime": {
            "title": "models.ReqPlatformUserTime",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                },
                "BeginTime": {
                    "description": "开始时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EndTime": {
                    "description": "结束时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Limit": {
                    "description": "获取数量 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "索引偏移 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserName": {
                    "description": "实名 ",
                    "type": "string"
                }
            }
        },
        "models.ReqProcDrawRecord": {
            "title": "models.ReqProcDrawRecord",
            "type": "object",
            "properties": {
                "During": {
                    "type": "integer",
                    "format": "int64"
                },
                "Number": {
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqPurchaseExport": {
            "title": "models.ReqPurchaseExport",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "投资起始日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EndTime": {
                    "description": "投资结束日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Limit": {
                    "description": "限制取多少条 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "起始偏移 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Platform": {
                    "description": "渠道id ",
                    "type": "integer",
                    "format": "int64"
                },
                "RegisterBeginTime": {
                    "description": "注册起始时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "RegisterEndTime": {
                    "description": "注册截至时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                }
            }
        },
        "models.ReqPurchaseRecord": {
            "title": "models.ReqPurchaseRecord",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "开始时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EndTime": {
                    "description": "结束时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Type": {
                    "description": "\"CURRENT\"活期 \"FIXED\"定期 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqPurchaseRecordId": {
            "title": "models.ReqPurchaseRecordId",
            "type": "object",
            "properties": {
                "PurchaseRecordId": {
                    "description": "购买记录id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqPurchaseRecordNew": {
            "title": "models.ReqPurchaseRecordNew",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                },
                "BeginTime": {
                    "description": "投资起始日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EndTime": {
                    "description": "投资结束日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Limit": {
                    "description": "限制取多少条 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "起始偏移 ",
                    "type": "integer",
                    "format": "int64"
                },
                "State": {
                    "description": "投资状态 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "uid ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserName": {
                    "description": "姓名 ",
                    "type": "string"
                }
            }
        },
        "models.ReqPurchaseRecordPoolId": {
            "title": "models.ReqPurchaseRecordPoolId",
            "type": "object",
            "properties": {
                "PurchaseRecordPoolId": {
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqPurchaseStatisticsRecord": {
            "title": "models.ReqPurchaseStatisticsRecord",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                },
                "BeginTime": {
                    "description": "起始日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "BidName": {
                    "description": "标名 ",
                    "type": "string"
                },
                "BidNo": {
                    "description": "标的编号 ",
                    "type": "string"
                },
                "EndTime": {
                    "description": "结束日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Limit": {
                    "description": "限制取多少条 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "起始偏移 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Platform": {
                    "description": "渠道id ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserName": {
                    "description": "用户真实姓名 ",
                    "type": "string"
                }
            }
        },
        "models.ReqPwd": {
            "title": "models.ReqPwd",
            "type": "object",
            "properties": {
                "Pwd": {
                    "description": "用户密码 ",
                    "type": "string"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "string"
                }
            }
        },
        "models.ReqRegiste": {
            "title": "models.ReqRegiste",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                },
                "DeviceNumber": {
                    "description": "设备号 ",
                    "type": "string"
                },
                "Pwd": {
                    "description": "密码 ",
                    "type": "string"
                },
                "Version": {
                    "description": "版本号 ",
                    "type": "string"
                }
            }
        },
        "models.ReqSetInviter": {
            "title": "models.ReqSetInviter",
            "type": "object",
            "properties": {
                "Inviter": {
                    "description": "用户邀请人的用户id ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqSetUserTag": {
            "title": "models.ReqSetUserTag",
            "type": "object",
            "properties": {
                "Key": {
                    "description": "用户标签键 ",
                    "type": "string"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                },
                "Value": {
                    "description": "用户标签值 ",
                    "type": "string"
                }
            }
        },
        "models.ReqUpdateBid": {
            "title": "models.ReqUpdateBid",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "开卖时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "BidDesc": {
                    "description": "标描述 ",
                    "type": "string"
                },
                "BidId": {
                    "description": "标id ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidInfoPath": {
                    "description": "标的详情地址 ",
                    "type": "string"
                },
                "BidMark": {
                    "description": "标的购买限制 ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidName": {
                    "description": "标名称 ",
                    "type": "string"
                },
                "BidRateShow": {
                    "description": "标利率特殊显示 ",
                    "type": "string"
                },
                "CheckDate": {
                    "description": "结算日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "DaysLimit": {
                    "description": "这是多少天标 ",
                    "type": "integer",
                    "format": "int64"
                },
                "EndTime": {
                    "description": "结束时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "MaxBuySize": {
                    "description": "最大购买笔数 ",
                    "type": "integer",
                    "format": "int64"
                },
                "MaxNumber": {
                    "description": "最大金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "MinNumber": {
                    "description": "起投金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillBidMark": {
                    "description": "标的购买限制 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillBidType": {
                    "description": "标计息类型 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillDaysLimit": {
                    "description": "标的物最大期限 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillRemainNumberMin": {
                    "description": "全部符合条件的标 累计小于最小续标金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "StepNumber": {
                    "description": "增幅金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "TotalNumber": {
                    "description": "标总金额 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqUpdateBidNew": {
            "title": "models.ReqUpdateBidNew",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "开卖时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "BidDesc": {
                    "description": "标描述 ",
                    "type": "string"
                },
                "BidId": {
                    "description": "标id ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidInfoPath": {
                    "description": "标的详情地址 ",
                    "type": "string"
                },
                "BidMark": {
                    "description": "标的购买限制 ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidName": {
                    "description": "标名称 ",
                    "type": "string"
                },
                "BidRateShow": {
                    "description": "标利率特殊显示 ",
                    "type": "string"
                },
                "EndTime": {
                    "description": "结束时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "MaxBuySize": {
                    "description": "最大购买笔数 ",
                    "type": "integer",
                    "format": "int64"
                },
                "MaxNumber": {
                    "description": "最大金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "MinNumber": {
                    "description": "起投金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillBidMark": {
                    "description": "标的购买限制 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillBidType": {
                    "description": "标计息类型 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillDaysLimit": {
                    "description": "标的物最大期限 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillRemainNumberMin": {
                    "description": "全部符合条件的标 累计小于最小续标金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "StepNumber": {
                    "description": "增幅金额 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqUpdateBidPool": {
            "title": "models.ReqUpdateBidPool",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "开卖时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "BidDesc": {
                    "description": "标描述 ",
                    "type": "string"
                },
                "BidId": {
                    "description": "标id ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidInfoPath": {
                    "description": "标的详情地址 ",
                    "type": "string"
                },
                "BidName": {
                    "description": "标名称 ",
                    "type": "string"
                },
                "CheckDate": {
                    "description": "结算日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "DaysLimit": {
                    "description": "这是多少天标 ",
                    "type": "integer",
                    "format": "int64"
                },
                "EndTime": {
                    "description": "结束时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "TotalNumber": {
                    "description": "标总金额 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqUpdateBidSatisfy": {
            "title": "models.ReqUpdateBidSatisfy",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "开卖时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "BidDesc": {
                    "description": "标描述 ",
                    "type": "string"
                },
                "BidId": {
                    "description": "标id ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidInfoPath": {
                    "description": "标的详情地址 ",
                    "type": "string"
                },
                "BidMark": {
                    "description": "标的购买限制 ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidName": {
                    "description": "标名称 ",
                    "type": "string"
                },
                "BidRateShow": {
                    "description": "标利率特殊显示 ",
                    "type": "string"
                },
                "CheckDate": {
                    "description": "结算日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "DaysLimit": {
                    "description": "这是多少天标 ",
                    "type": "integer",
                    "format": "int64"
                },
                "EndTime": {
                    "description": "结束时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "MaxBuySize": {
                    "description": "最大购买笔数 ",
                    "type": "integer",
                    "format": "int64"
                },
                "MaxNumber": {
                    "description": "最大金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "MinNumber": {
                    "description": "起投金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillBidMark": {
                    "description": "标的购买限制 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillBidType": {
                    "description": "标计息类型 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillDaysLimit": {
                    "description": "标的物最大期限 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillRemainNumberMin": {
                    "description": "全部符合条件的标 累计小于最小续标金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "StepNumber": {
                    "description": "增幅金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "TotalNumber": {
                    "description": "标总金额 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqUpdatePlatform": {
            "title": "models.ReqUpdatePlatform",
            "type": "object",
            "properties": {
                "Id": {
                    "description": "id ",
                    "type": "integer",
                    "format": "int64"
                },
                "Name": {
                    "description": "渠道名称 ",
                    "type": "string"
                },
                "PackCode": {
                    "description": "包名 ",
                    "type": "string"
                },
                "PlatformId": {
                    "description": "渠道id ",
                    "type": "integer",
                    "format": "int64"
                },
                "State": {
                    "description": "状态 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqUserAgent": {
            "title": "models.ReqUserAgent",
            "type": "object",
            "properties": {
                "Agent": {
                    "description": "useragent ",
                    "type": "string"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqUserAndPage": {
            "title": "models.ReqUserAndPage",
            "type": "object",
            "properties": {
                "Limit": {
                    "description": "获取数量 ",
                    "type": "integer",
                    "format": "int64"
                },
                "NumberBegin": {
                    "description": "最小金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "NumberEnd": {
                    "description": "最大金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "索引偏移 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Query": {
                    "description": "查询 ",
                    "type": "string"
                },
                "TimeBegin": {
                    "description": "开始时间 ",
                    "type": "string"
                },
                "TimeEnd": {
                    "description": "结束时间 ",
                    "type": "string"
                }
            }
        },
        "models.ReqUserBidState": {
            "title": "models.ReqUserBidState",
            "type": "object",
            "properties": {
                "BidState": {
                    "description": "募集状态 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqUserCouponId": {
            "title": "models.ReqUserCouponId",
            "type": "object",
            "properties": {
                "UserCouponId": {
                    "description": "用户的优惠券id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqUserId": {
            "title": "models.ReqUserId",
            "type": "object",
            "properties": {
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqUserIdOffsetLimit": {
            "title": "models.ReqUserIdOffsetLimit",
            "type": "object",
            "properties": {
                "Limit": {
                    "description": "获取数量 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "索引偏移 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqUserIdcard": {
            "title": "models.ReqUserIdcard",
            "type": "object",
            "properties": {
                "Idcard": {
                    "description": "身份证 ",
                    "type": "string"
                }
            }
        },
        "models.ReqUserPlatform": {
            "title": "models.ReqUserPlatform",
            "type": "object",
            "properties": {
                "Platform": {
                    "description": "平台id ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqUserState": {
            "title": "models.ReqUserState",
            "type": "object",
            "properties": {
                "State": {
                    "description": "状态 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqUserTime": {
            "title": "models.ReqUserTime",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                },
                "BeginTime": {
                    "description": "开始时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EndTime": {
                    "description": "结束时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Limit": {
                    "description": "限制 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "偏移 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserName": {
                    "description": "实名 ",
                    "type": "string"
                }
            }
        },
        "models.ReqUserTimeRange": {
            "title": "models.ReqUserTimeRange",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "开始时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EndTime": {
                    "description": "结束时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqUserTimeRangeOffsetLimit": {
            "title": "models.ReqUserTimeRangeOffsetLimit",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "开始时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EndTime": {
                    "description": "结束时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Limit": {
                    "description": "限制 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "偏移 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqWithdraw": {
            "title": "models.ReqWithdraw",
            "type": "object",
            "properties": {
                "BankcardId": {
                    "description": "银行卡id ",
                    "type": "integer",
                    "format": "int64"
                },
                "Number": {
                    "description": "金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqWithdrawApply": {
            "title": "models.ReqWithdrawApply",
            "type": "object",
            "properties": {
                "OrderNo": {
                    "description": "订单号 ",
                    "type": "string"
                },
                "WithdrawId": {
                    "description": "提现id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqWithdrawCount": {
            "title": "models.ReqWithdrawCount",
            "type": "object",
            "properties": {
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                },
                "WithdrawTime": {
                    "description": "提现时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                }
            }
        },
        "models.ReqWithdrawRecord": {
            "title": "models.ReqWithdrawRecord",
            "type": "object",
            "properties": {
                "Limit": {
                    "description": "获取数量 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "索引偏移 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqWithdrawState": {
            "title": "models.ReqWithdrawState",
            "type": "object",
            "properties": {
                "Mark": {
                    "description": "标记 ",
                    "type": "string"
                },
                "State": {
                    "description": "状态 ",
                    "type": "integer",
                    "format": "int64"
                },
                "WithdrawId": {
                    "description": "提现id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReqWithdrawUser": {
            "title": "models.ReqWithdrawUser",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "起始时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EndTime": {
                    "description": "结束时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Limit": {
                    "description": "条数 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Offset": {
                    "description": "偏移 ",
                    "type": "integer",
                    "format": "int64"
                },
                "State": {
                    "description": "状态 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.RespBankcard": {
            "title": "models.RespBankcard",
            "type": "object",
            "properties": {
                "BankCode": {
                    "type": "string"
                },
                "BankDesc": {
                    "type": "string"
                },
                "BankName": {
                    "type": "string"
                },
                "BindId": {
                    "type": "string"
                },
                "Card": {
                    "type": "string"
                },
                "CreateTime": {
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "DayLimit": {
                    "type": "string"
                },
                "Id": {
                    "type": "integer",
                    "format": "int64"
                },
                "PayWay": {
                    "type": "integer",
                    "format": "int64"
                },
                "Phone": {
                    "type": "string"
                },
                "SingleLimit": {
                    "type": "string"
                },
                "State": {
                    "type": "integer",
                    "format": "int64"
                },
                "TotalIn": {
                    "type": "integer",
                    "format": "int64"
                },
                "TotalOut": {
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.RespBids": {
            "title": "models.RespBids",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "开卖时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "BidDesc": {
                    "description": "标描述 ",
                    "type": "string"
                },
                "BidId": {
                    "description": "标的id ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidInfoPath": {
                    "description": "标详情路径 ",
                    "type": "string"
                },
                "BidMark": {
                    "description": "标的购买限制 ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidName": {
                    "description": "标名称 ",
                    "type": "string"
                },
                "BidNo": {
                    "description": "标订单号 ",
                    "type": "string"
                },
                "BidRate": {
                    "description": "标的物利率 ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidRateBase": {
                    "description": "基础利率 ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidRateRaise": {
                    "description": "加息利率 ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidRateShow": {
                    "description": "标利率特殊显示 ",
                    "type": "string"
                },
                "BidType": {
                    "description": "标计息类型 标记活期定期 ",
                    "type": "integer",
                    "format": "int64"
                },
                "CheckDate": {
                    "description": "结算日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Days": {
                    "description": "实际经过的天数 ",
                    "type": "integer",
                    "format": "int64"
                },
                "DaysLimit": {
                    "description": "这是多少天标 ",
                    "type": "integer",
                    "format": "int64"
                },
                "EndTime": {
                    "description": "结束时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Id": {
                    "description": "标的id ",
                    "type": "integer",
                    "format": "int64"
                },
                "IncomeDate": {
                    "description": "起息时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "IsRank": {
                    "description": "是计入rank的 ",
                    "type": "integer",
                    "format": "int64"
                },
                "IsVip": {
                    "description": "是计入vip的 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Mark": {
                    "description": "标记 ",
                    "type": "string"
                },
                "MaxBuySize": {
                    "description": "最大购买笔数 ",
                    "type": "integer",
                    "format": "int64"
                },
                "MaxNumber": {
                    "description": "最大金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "MinNumber": {
                    "description": "起投金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillBidMark": {
                    "description": "标的购买限制 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillBidType": {
                    "description": "标计息类型 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillDaysLimit": {
                    "description": "标的物最大期限 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillRemainNumberMin": {
                    "description": "全部符合条件的标 累计小于最小续标金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RemainNumber": {
                    "description": "标剩余金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "StepNumber": {
                    "description": "增幅金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "TotalNumber": {
                    "description": "标总金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UpdateTime": {
                    "description": "更新时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                }
            }
        },
        "models.RespChargeRecordStatistics": {
            "title": "models.RespChargeRecordStatistics",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                },
                "BankCode": {
                    "description": "银行编码 ",
                    "type": "string"
                },
                "BankName": {
                    "description": "银行名 ",
                    "type": "string"
                },
                "Card": {
                    "description": "卡号 ",
                    "type": "string"
                },
                "CreateTime": {
                    "description": "充值时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Number": {
                    "description": "充值金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "PayNo": {
                    "description": "订单号 ",
                    "type": "string"
                },
                "PayWay": {
                    "description": "支付渠道 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserName": {
                    "description": "实名 ",
                    "type": "string"
                }
            }
        },
        "models.RespDrawRecord": {
            "title": "models.RespDrawRecord",
            "type": "object",
            "properties": {
                "Account": {
                    "type": "string"
                },
                "Card": {
                    "type": "string"
                },
                "CheckState": {
                    "type": "integer",
                    "format": "int64"
                },
                "CheckTime": {
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "CreateTime": {
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "DrawNumber": {
                    "type": "integer",
                    "format": "int64"
                },
                "Id": {
                    "type": "integer",
                    "format": "int64"
                },
                "State": {
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "type": "integer",
                    "format": "int64"
                },
                "UserType": {
                    "type": "string"
                }
            }
        },
        "models.RespEstAddPurchase": {
            "title": "models.RespEstAddPurchase",
            "type": "object",
            "properties": {
                "BidId": {
                    "description": "标的id ",
                    "type": "integer",
                    "format": "int64"
                },
                "CouponRules": {
                    "description": "购买时使用的券 ",
                    "type": "string"
                },
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EarningsNumber": {
                    "description": "预期收益 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Id": {
                    "type": "integer",
                    "format": "int64"
                },
                "Number": {
                    "description": "购买金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "PayNumber": {
                    "description": "支付金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "PurchaseRecordPoolId": {
                    "description": "购买记录池id ",
                    "type": "integer",
                    "format": "int64"
                },
                "RemainNumber": {
                    "description": "剩余金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UpdateTime": {
                    "description": "更新时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.RespExperienceHistory": {
            "title": "models.RespExperienceHistory",
            "type": "object",
            "properties": {
                "EarningsNumber": {
                    "type": "integer",
                    "format": "int64"
                },
                "Number": {
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.RespFeedbackRecord": {
            "title": "models.RespFeedbackRecord",
            "type": "object",
            "properties": {
                "Account": {
                    "type": "string"
                },
                "Content": {
                    "type": "string"
                },
                "CreateTime": {
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Id": {
                    "type": "integer",
                    "format": "int64"
                },
                "UserName": {
                    "type": "string"
                }
            }
        },
        "models.RespFullBid": {
            "title": "models.RespFullBid",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "开卖时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "BidDesc": {
                    "description": "标描述 ",
                    "type": "string"
                },
                "BidInfoPath": {
                    "description": "标详情路径 ",
                    "type": "string"
                },
                "BidMark": {
                    "description": "标的物分类 ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidName": {
                    "description": "标名称 ",
                    "type": "string"
                },
                "BidNo": {
                    "description": "标订单号 ",
                    "type": "string"
                },
                "BidRateShow": {
                    "description": "标利率特殊显示 ",
                    "type": "string"
                },
                "BidRule": {
                    "description": "限制规则 ",
                    "type": "string"
                },
                "BidType": {
                    "description": "标计息类型 标记活期定期 ",
                    "type": "integer",
                    "format": "int64"
                },
                "CheckDate": {
                    "description": "结算日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "DaysLimit": {
                    "description": "这是多少天标 ",
                    "type": "integer",
                    "format": "int64"
                },
                "EndTime": {
                    "description": "结束时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Id": {
                    "description": "标的id ",
                    "type": "integer",
                    "format": "int64"
                },
                "IncomeDate": {
                    "description": "起息时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "RemainNumber": {
                    "description": "标剩余金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "TotalNumber": {
                    "description": "标总金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UpdateTime": {
                    "description": "修改时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                }
            }
        },
        "models.RespGetRanking": {
            "title": "models.RespGetRanking",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                },
                "NickName": {
                    "description": "昵称 ",
                    "type": "string"
                },
                "Rank": {
                    "description": "分数 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserName": {
                    "description": "真实姓名 ",
                    "type": "string"
                }
            }
        },
        "models.RespGetUserCouponNew": {
            "title": "models.RespGetUserCouponNew",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "生效时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "CouponDesc": {
                    "description": "优惠券信息 ",
                    "type": "string"
                },
                "CouponId": {
                    "description": "优惠券类型id ",
                    "type": "integer",
                    "format": "int64"
                },
                "CouponName": {
                    "description": "优惠券描述 ",
                    "type": "string"
                },
                "CouponRule": {
                    "description": "优惠券规则 ",
                    "type": "string"
                },
                "CouponType": {
                    "description": "优惠券类型 0 | CouponTypeDiscount | CouponTypeRaise ",
                    "type": "integer",
                    "format": "int64"
                },
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EndTime": {
                    "description": "失效时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Logo": {
                    "description": "logo ",
                    "type": "string"
                },
                "State": {
                    "description": "优惠券状态 可选 0 | UserCouponStateUnused | UserCouponStateUsed | UserCouponStateOverdue ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserCouponId": {
                    "description": "优惠券id ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.RespMessage": {
            "title": "models.RespMessage",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "显示时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "CreateTime": {
                    "description": "发布时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "MessageId": {
                    "description": "id ",
                    "type": "integer",
                    "format": "int64"
                },
                "Msg": {
                    "description": "内容 ",
                    "type": "string"
                },
                "MsgType": {
                    "description": "消息类型 ",
                    "type": "integer",
                    "format": "int64"
                },
                "ReadCount": {
                    "description": "已读数量 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Title": {
                    "description": "标题 ",
                    "type": "string"
                },
                "TotalCount": {
                    "description": "总数 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.RespPurchaseExport": {
            "title": "models.RespPurchaseExport",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                },
                "Bankcard": {
                    "description": "用户银行卡号 ",
                    "type": "string"
                },
                "BidName": {
                    "description": "标名 ",
                    "type": "string"
                },
                "CreateTime": {
                    "description": "购买时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "DaysLimit": {
                    "description": "期限天数 ",
                    "type": "integer",
                    "format": "int64"
                },
                "IdCard": {
                    "description": "身份证号 ",
                    "type": "string"
                },
                "Inviter": {
                    "description": "邀请人id ",
                    "type": "integer",
                    "format": "int64"
                },
                "InviterAccount": {
                    "description": "邀请人账号 ",
                    "type": "string"
                },
                "InviterName": {
                    "description": "邀请人实名 ",
                    "type": "string"
                },
                "Number": {
                    "description": "投资金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "PayNumber": {
                    "description": "实付金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Platform": {
                    "description": "渠道 ",
                    "type": "integer",
                    "format": "int64"
                },
                "PurchaseCount": {
                    "description": "购买次数 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RegisterTime": {
                    "description": "注册时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserName": {
                    "description": "实名 ",
                    "type": "string"
                }
            }
        },
        "models.RespPurchaseFullCount": {
            "title": "models.RespPurchaseFullCount",
            "type": "object",
            "properties": {
                "EveryFullCount": {
                    "type": "integer",
                    "format": "int64"
                },
                "FullCount": {
                    "type": "integer",
                    "format": "int64"
                },
                "PurchaseNumber": {
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.RespPurchaseRecord": {
            "title": "models.RespPurchaseRecord",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "开卖时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "BidDesc": {
                    "description": "标描述 ",
                    "type": "string"
                },
                "BidId": {
                    "description": "标的id ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidInfoPath": {
                    "description": "标详情路径 ",
                    "type": "string"
                },
                "BidMark": {
                    "description": "标的购买限制 ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidName": {
                    "description": "标名称 ",
                    "type": "string"
                },
                "BidNo": {
                    "description": "标订单号 ",
                    "type": "string"
                },
                "BidRate": {
                    "description": "标的物利率 ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidRateBase": {
                    "description": "基础利率 ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidRateRaise": {
                    "description": "加息利率 ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidRateShow": {
                    "description": "标利率特殊显示 ",
                    "type": "string"
                },
                "BidType": {
                    "description": "标计息类型 标记活期定期 ",
                    "type": "integer",
                    "format": "int64"
                },
                "CheckDate": {
                    "description": "结算日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "CouponMsg": {
                    "description": "优惠券使用描述 ",
                    "type": "string"
                },
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Days": {
                    "description": "实际经过的天数 ",
                    "type": "integer",
                    "format": "int64"
                },
                "DaysLimit": {
                    "description": "这是多少天标 ",
                    "type": "integer",
                    "format": "int64"
                },
                "DrawNumber": {
                    "description": "持有金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "EarningsNumber": {
                    "description": "购买金额的收益 ",
                    "type": "integer",
                    "format": "int64"
                },
                "EndTime": {
                    "description": "结束时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "ExpireDays": {
                    "description": "预计天数 废弃 ",
                    "type": "integer",
                    "format": "int64"
                },
                "ExpireEarningsNumber": {
                    "description": "预计到期收益 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Id": {
                    "description": "标的id ",
                    "type": "integer",
                    "format": "int64"
                },
                "IncomeDate": {
                    "description": "起息时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "IsIncome": {
                    "description": "是否起息 ",
                    "type": "integer",
                    "format": "int64"
                },
                "IsRank": {
                    "description": "是计入rank的 ",
                    "type": "integer",
                    "format": "int64"
                },
                "IsVip": {
                    "description": "是计入vip的 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Mark": {
                    "description": "标记 ",
                    "type": "string"
                },
                "MaxBuySize": {
                    "description": "最大购买笔数 ",
                    "type": "integer",
                    "format": "int64"
                },
                "MaxNumber": {
                    "description": "最大金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "MinNumber": {
                    "description": "起投金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Number": {
                    "description": "金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "PayNumber": {
                    "description": "支付金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "PurchaseRecordId": {
                    "description": "购买记录id ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillBidMark": {
                    "description": "标的购买限制 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillBidType": {
                    "description": "标计息类型 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillDaysLimit": {
                    "description": "标的物最大期限 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RefillRemainNumberMin": {
                    "description": "全部符合条件的标 累计小于最小续标金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RemainNumber": {
                    "description": "标剩余金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "StepNumber": {
                    "description": "增幅金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "TotalNumber": {
                    "description": "标总金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UpdateTime": {
                    "description": "更新时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                },
                "YesterEarningsNumber": {
                    "description": "昨日收益 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.RespPurchaseRecordNew": {
            "title": "models.RespPurchaseRecordNew",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                },
                "BeginTime": {
                    "description": "标起售时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "BidId": {
                    "description": "标id ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidName": {
                    "description": "标名 ",
                    "type": "string"
                },
                "BidRateShow": {
                    "description": "标利率展示 ",
                    "type": "string"
                },
                "CheckDate": {
                    "description": "回款时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "CouponRules": {
                    "description": "使用的券 ",
                    "type": "string"
                },
                "CreateTime": {
                    "description": "购买时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "DaysLimit": {
                    "description": "期限天数 ",
                    "type": "integer",
                    "format": "int64"
                },
                "EndTime": {
                    "description": "标结束时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Number": {
                    "description": "投资金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "PayNumber": {
                    "description": "实付金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "PurchaseRecordId": {
                    "description": "投资记录id ",
                    "type": "integer",
                    "format": "int64"
                },
                "PurchaseRecordPoolId": {
                    "description": "活期池id ",
                    "type": "integer",
                    "format": "int64"
                },
                "PurchaseRemainNumber": {
                    "description": "活期剩余金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Rate": {
                    "description": "标的利率 ",
                    "type": "integer",
                    "format": "int64"
                },
                "RemainNumber": {
                    "description": "标剩余金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserName": {
                    "description": "实名 ",
                    "type": "string"
                }
            }
        },
        "models.RespPurchaseRecordsSum": {
            "title": "models.RespPurchaseRecordsSum",
            "type": "object",
            "properties": {
                "LoansCount": {
                    "description": "当前还没回款的笔数 ",
                    "type": "integer",
                    "format": "int64"
                },
                "LoansNumber": {
                    "description": "当前还没回款的金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "ReceiptCount": {
                    "description": "当前已经回款的笔数 ",
                    "type": "integer",
                    "format": "int64"
                },
                "ReceiptNumber": {
                    "description": "当前已经回款的金额 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.RespPurchaseTotal": {
            "title": "models.RespPurchaseTotal",
            "type": "object",
            "properties": {
                "Number": {
                    "description": "投资金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "PayNumber": {
                    "description": "实付 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.RespReturnedRecordInfo": {
            "title": "models.RespReturnedRecordInfo",
            "type": "object",
            "properties": {
                "BidId": {
                    "description": "标id ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidName": {
                    "description": "标名 ",
                    "type": "string"
                },
                "CheckDate": {
                    "description": "还款日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Days": {
                    "description": "存款天数 废除 ",
                    "type": "integer",
                    "format": "int64"
                },
                "DaysLimit": {
                    "description": "标的物总天数 ",
                    "type": "integer",
                    "format": "int64"
                },
                "DrawNumber": {
                    "description": "回款本金 ",
                    "type": "integer",
                    "format": "int64"
                },
                "EarningsNumber": {
                    "description": "回款利息 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Id": {
                    "description": "废除 ",
                    "type": "integer",
                    "format": "int64"
                },
                "PurchaseRecordId": {
                    "description": "购买记录id ",
                    "type": "integer",
                    "format": "int64"
                },
                "PurchaseRecordPoolId": {
                    "description": "购买记录池id ",
                    "type": "integer",
                    "format": "int64"
                },
                "Stages": {
                    "description": "第几期 ",
                    "type": "integer",
                    "format": "int64"
                },
                "State": {
                    "description": "状态 ",
                    "type": "integer",
                    "format": "int64"
                },
                "TotalStages": {
                    "description": "共几期 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UpdateTime": {
                    "description": "更新时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                },
                "YesterEarningsNumber": {
                    "description": "昨天回款利息 废除 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.RespReturnedSum": {
            "title": "models.RespReturnedSum",
            "type": "object",
            "properties": {
                "DrawNumber": {
                    "description": "历史投资金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "EarningsNumber": {
                    "description": "历史投资收益 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.RespStatisticsPurchaseInfo": {
            "title": "models.RespStatisticsPurchaseInfo",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                },
                "BidId": {
                    "description": "标id ",
                    "type": "integer",
                    "format": "int64"
                },
                "BidName": {
                    "description": "标名称 ",
                    "type": "string"
                },
                "BidNo": {
                    "description": "标编号 ",
                    "type": "string"
                },
                "BidRateShow": {
                    "description": "标利率 ",
                    "type": "string"
                },
                "CheckDate": {
                    "description": "回款日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "CouponName": {
                    "description": "券名称 ",
                    "type": "string"
                },
                "CouponRules": {
                    "description": "券规则 ",
                    "type": "string"
                },
                "CreateTime": {
                    "description": "投资时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "DaysLimit": {
                    "description": "标期限 ",
                    "type": "integer",
                    "format": "int64"
                },
                "EndTime": {
                    "description": "止卖时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Inviter": {
                    "description": "邀请人id ",
                    "type": "integer",
                    "format": "int64"
                },
                "InviterAccount": {
                    "description": "邀请人账号 ",
                    "type": "string"
                },
                "Number": {
                    "description": "投资金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "PayNumber": {
                    "description": "实付金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Platform": {
                    "description": "渠道 ",
                    "type": "integer",
                    "format": "int64"
                },
                "PlatformName": {
                    "description": "渠道名称 ",
                    "type": "string"
                },
                "RemainNumber": {
                    "description": "标剩余金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "TotalNumber": {
                    "description": "标总金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserName": {
                    "description": "姓名 ",
                    "type": "string"
                }
            }
        },
        "models.RespToday": {
            "title": "models.RespToday",
            "type": "object",
            "properties": {
                "Number": {
                    "description": "购买金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "PayNumber": {
                    "description": "支付金额 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.RespUser": {
            "title": "models.RespUser",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                },
                "AccountType": {
                    "description": "账号类型 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Agent": {
                    "description": "注册设备信息 ",
                    "type": "string"
                },
                "CreateTime": {
                    "description": "注册时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Gender": {
                    "description": "性别 ",
                    "type": "string"
                },
                "IdCard": {
                    "description": "身份证号 ",
                    "type": "string"
                },
                "Inviter": {
                    "description": "邀请人id ",
                    "type": "integer",
                    "format": "int64"
                },
                "NickName": {
                    "description": "昵称 ",
                    "type": "string"
                },
                "Platform": {
                    "description": "渠道 ",
                    "type": "string"
                },
                "PlatformName": {
                    "description": "渠道名称 ",
                    "type": "string"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserName": {
                    "description": "实名 ",
                    "type": "string"
                }
            }
        },
        "models.RespUserAssets": {
            "title": "models.RespUserAssets",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                },
                "FreezeNumber": {
                    "description": "冻结金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Number": {
                    "description": "可用余额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "TotalEarningsNumber": {
                    "description": "总回款收益 ",
                    "type": "integer",
                    "format": "int64"
                },
                "TotalNumber": {
                    "description": "总投资金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "TotalPayNumber": {
                    "description": "总实付金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "TotalReturnNumber": {
                    "description": "总回款本金 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserName": {
                    "description": "实名 ",
                    "type": "string"
                }
            }
        },
        "models.RespUserBalanceChange": {
            "title": "models.RespUserBalanceChange",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                },
                "CreateTime": {
                    "description": "时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Msg": {
                    "description": "交易信息 ",
                    "type": "string"
                },
                "Number": {
                    "description": "变更金额 ",
                    "type": "string"
                },
                "UserName": {
                    "description": "实名 ",
                    "type": "string"
                }
            }
        },
        "models.RespUserBalanceChangeRecord": {
            "title": "models.RespUserBalanceChangeRecord",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                },
                "AfterNumber": {
                    "description": "变更后的金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "CreateTime": {
                    "description": "时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Msg": {
                    "description": "交易信息 ",
                    "type": "string"
                },
                "Number": {
                    "description": "变更金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserName": {
                    "description": "实名 ",
                    "type": "string"
                }
            }
        },
        "models.RespUserForSrarch": {
            "title": "models.RespUserForSrarch",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                },
                "CreateTime": {
                    "description": "注册时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "IdCard": {
                    "description": "身份证号 ",
                    "type": "string"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserName": {
                    "description": "实名 ",
                    "type": "string"
                }
            }
        },
        "models.RespUserInfoWithAccount": {
            "title": "models.RespUserInfoWithAccount",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                },
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Gender": {
                    "description": "性别 ",
                    "type": "string"
                },
                "IdCard": {
                    "description": "身份证号 ",
                    "type": "string"
                },
                "Inviter": {
                    "description": "邀请人 id ",
                    "type": "integer",
                    "format": "int64"
                },
                "NickName": {
                    "description": "用户名 ",
                    "type": "string"
                },
                "Pic": {
                    "description": "头像地址 ",
                    "type": "string"
                },
                "UserId": {
                    "description": "用户信息id ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserName": {
                    "description": "真实姓名 ",
                    "type": "string"
                }
            }
        },
        "models.RespUserPurchaseStatement": {
            "title": "models.RespUserPurchaseStatement",
            "type": "object",
            "properties": {
                "NewCount": {
                    "description": "新手标投资次数 ",
                    "type": "integer",
                    "format": "int64"
                },
                "NewNumber": {
                    "description": "新手标投资金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "TotalCount": {
                    "description": "总投资次数 ",
                    "type": "integer",
                    "format": "int64"
                },
                "TotalNumber": {
                    "description": "总投资金额 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.RespUsers": {
            "title": "models.RespUsers",
            "type": "object"
        },
        "models.RespWithdrawRecord": {
            "title": "models.RespWithdrawRecord",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                },
                "BankCode": {
                    "description": "银行编号 ",
                    "type": "string"
                },
                "BankName": {
                    "description": "银行名称 ",
                    "type": "string"
                },
                "BankcardId": {
                    "description": "银行卡id ",
                    "type": "integer",
                    "format": "int64"
                },
                "Card": {
                    "description": "卡号 ",
                    "type": "string"
                },
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "DrawNumber": {
                    "description": "提现金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Id": {
                    "description": "提现记录id ",
                    "type": "integer",
                    "format": "int64"
                },
                "State": {
                    "description": "状态 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UpdateTime": {
                    "description": "更新时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserName": {
                    "description": "姓名 ",
                    "type": "string"
                }
            }
        },
        "models.RespWithdrawRecordStatistics": {
            "title": "models.RespWithdrawRecordStatistics",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                },
                "BankCode": {
                    "description": "银行编码 ",
                    "type": "string"
                },
                "BankName": {
                    "description": "银行名称 ",
                    "type": "string"
                },
                "Card": {
                    "description": "卡号 ",
                    "type": "string"
                },
                "CreateTime": {
                    "description": "提现申请时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "DrawNumber": {
                    "description": "提现金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "State": {
                    "description": "提现状态 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UpdateTime": {
                    "description": "状态更新时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserName": {
                    "description": "实名 ",
                    "type": "string"
                },
                "WithdrawId": {
                    "description": "提现id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ResqBalanceChange": {
            "title": "models.ResqBalanceChange",
            "type": "object",
            "properties": {
                "CreateTime": {
                    "description": "变动的时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Msg": {
                    "description": "变动描述 ",
                    "type": "string"
                },
                "Number": {
                    "description": "变动的金额 ",
                    "type": "string"
                }
            }
        },
        "models.ResqReturnRecordInfo": {
            "title": "models.ResqReturnRecordInfo",
            "type": "object",
            "properties": {
                "BidName": {
                    "type": "string"
                },
                "CreateTime": {
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "DrawNumber": {
                    "type": "integer",
                    "format": "int64"
                },
                "EarningsNumber": {
                    "type": "integer",
                    "format": "int64"
                },
                "UserName": {
                    "type": "string"
                }
            }
        },
        "models.ResqStatisticsChargeInfo": {
            "title": "models.ResqStatisticsChargeInfo",
            "type": "object",
            "properties": {
                "ChargeNumber": {
                    "type": "integer",
                    "format": "int64"
                },
                "DateTime": {
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                }
            }
        },
        "models.ResqStatisticsPurchaseInfo": {
            "title": "models.ResqStatisticsPurchaseInfo",
            "type": "object",
            "properties": {
                "DateTime": {
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "DiscountNumber": {
                    "type": "integer",
                    "format": "int64"
                },
                "Number": {
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ResqStatisticsRegisterInfo": {
            "title": "models.ResqStatisticsRegisterInfo",
            "type": "object",
            "properties": {
                "DateTime": {
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "RegisterNumber": {
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ResqStatisticsReturnInfo": {
            "title": "models.ResqStatisticsReturnInfo",
            "type": "object",
            "properties": {
                "DateTime": {
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "DrawNumber": {
                    "type": "integer",
                    "format": "int64"
                },
                "EarningsNumber": {
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ResqStatisticsVerifyInfo": {
            "title": "models.ResqStatisticsVerifyInfo",
            "type": "object",
            "properties": {
                "DateTime": {
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "VerifyNumber": {
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ResqUserChargeRecord": {
            "title": "models.ResqUserChargeRecord",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "电话号码 ",
                    "type": "string"
                },
                "BankcardId": {
                    "description": "银行卡id ",
                    "type": "integer",
                    "format": "int64"
                },
                "Card": {
                    "description": "银行卡号 ",
                    "type": "string"
                },
                "CreateTime": {
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Id": {
                    "type": "integer",
                    "format": "int64"
                },
                "Number": {
                    "description": "支付金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "PayNo": {
                    "description": "订单号 ",
                    "type": "string"
                },
                "PayWay": {
                    "description": "支付方式 ",
                    "type": "integer",
                    "format": "int64"
                },
                "TpOrderNo": {
                    "description": "第三方订单号 ",
                    "type": "string"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserName": {
                    "description": "真实姓名 ",
                    "type": "string"
                }
            }
        },
        "models.ResqUserRegisteInfo": {
            "title": "models.ResqUserRegisteInfo",
            "type": "object",
            "properties": {
                "DateTime": {
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "RegisterNumber": {
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.Returned": {
            "title": "models.Returned",
            "type": "object",
            "properties": {
                "BidId": {
                    "description": "标的id ",
                    "type": "integer",
                    "format": "int64"
                },
                "CheckDate": {
                    "description": "还款日期 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "DrawNumber": {
                    "description": "还款本金 ",
                    "type": "integer",
                    "format": "int64"
                },
                "EarningsNumber": {
                    "description": "还款利息 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Id": {
                    "type": "integer",
                    "format": "int64"
                },
                "PurchaseRecordId": {
                    "description": "购买记录id ",
                    "type": "integer",
                    "format": "int64"
                },
                "Stages": {
                    "description": "第几期 ",
                    "type": "integer",
                    "format": "int64"
                },
                "State": {
                    "description": "状态 ",
                    "type": "integer",
                    "format": "int64"
                },
                "TotalStages": {
                    "description": "总的几期 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UpdateTime": {
                    "description": "更新时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.ReturnedRecord": {
            "title": "models.ReturnedRecord",
            "type": "object",
            "properties": {
                "BidId": {
                    "description": "标 id ",
                    "type": "integer",
                    "format": "int64"
                },
                "CreateTime": {
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Days": {
                    "description": "存款天数 ",
                    "type": "integer",
                    "format": "int64"
                },
                "DrawNumber": {
                    "description": "回款本金 ",
                    "type": "integer",
                    "format": "int64"
                },
                "EarningsNumber": {
                    "description": "回款利息 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Id": {
                    "type": "integer",
                    "format": "int64"
                },
                "PurchaseRecordId": {
                    "description": "购买记录id ",
                    "type": "integer",
                    "format": "int64"
                },
                "PurchaseRecordPoolId": {
                    "description": "购买记录池id ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户 id ",
                    "type": "integer",
                    "format": "int64"
                },
                "YesterEarningsNumber": {
                    "description": "昨天回款利息 ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.UserAccount": {
            "title": "models.UserAccount",
            "type": "object",
            "properties": {
                "Account": {
                    "description": "账号 ",
                    "type": "string"
                },
                "AccountType": {
                    "description": "账号类型 ",
                    "type": "integer",
                    "format": "int64"
                },
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Id": {
                    "description": "账用户号id ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.UserCoupon": {
            "title": "models.UserCoupon",
            "type": "object",
            "properties": {
                "BeginTime": {
                    "description": "生效时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "CouponId": {
                    "description": "优惠券id ",
                    "type": "integer",
                    "format": "int64"
                },
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "EndTime": {
                    "description": "失效时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Id": {
                    "description": "用户优惠券id ",
                    "type": "integer",
                    "format": "int64"
                },
                "State": {
                    "description": "使用状态 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "models.UserInfos": {
            "title": "models.UserInfos",
            "type": "object",
            "properties": {
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "Gender": {
                    "description": "性别 ",
                    "type": "string"
                },
                "IdCard": {
                    "description": "身份证号 ",
                    "type": "string"
                },
                "Inviter": {
                    "description": "邀请人 id ",
                    "type": "integer",
                    "format": "int64"
                },
                "NickName": {
                    "description": "用户名 ",
                    "type": "string"
                },
                "Pic": {
                    "description": "头像地址 ",
                    "type": "string"
                },
                "UserId": {
                    "description": "用户信息id ",
                    "type": "integer",
                    "format": "int64"
                },
                "UserName": {
                    "description": "真实姓名 ",
                    "type": "string"
                }
            }
        },
        "models.UserNameReq": {
            "title": "models.UserNameReq",
            "type": "object"
        },
        "models.UserPwd": {
            "title": "models.UserPwd",
            "type": "object"
        },
        "models.Users": {
            "title": "models.Users",
            "type": "object",
            "properties": {
                "Agent": {
                    "description": "useragent ",
                    "type": "string"
                },
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "DeviceNumber": {
                    "description": "设备驱动号 ",
                    "type": "string"
                },
                "Id": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                },
                "PayPwd": {
                    "description": "支付密码 ",
                    "type": "string"
                },
                "PaySalt": {
                    "description": "支付密码加密矢量 ",
                    "type": "string"
                },
                "Platform": {
                    "description": "注册来源 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Pwd": {
                    "description": "密码 ",
                    "type": "string"
                },
                "Salt": {
                    "description": "加密矢量 ",
                    "type": "string"
                },
                "State": {
                    "description": "用户状态 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Version": {
                    "description": "版本号 ",
                    "type": "string"
                }
            }
        },
        "models.WithdrawRecord": {
            "title": "models.WithdrawRecord",
            "type": "object",
            "properties": {
                "BankcardId": {
                    "description": "银行卡id ",
                    "type": "integer",
                    "format": "int64"
                },
                "CreateTime": {
                    "description": "创建时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "DrawNumber": {
                    "description": "提现金额 ",
                    "type": "integer",
                    "format": "int64"
                },
                "Id": {
                    "description": "提现记录id ",
                    "type": "integer",
                    "format": "int64"
                },
                "State": {
                    "description": "提现状态 ",
                    "type": "integer",
                    "format": "int64"
                },
                "UpdateTime": {
                    "description": "更新时间 ",
                    "type": "string",
                    "example": "2006-01-02T15:04:05+08:00"
                },
                "UserId": {
                    "description": "用户id ",
                    "type": "integer",
                    "format": "int64"
                }
            }
        }
    },
    "tags": [
        {
            "name": "/balance",
            "description": "余额相关"
        },
        {
            "name": "/bankcard",
            "description": "银行卡相关"
        },
        {
            "name": "/banner",
            "description": "轮播图相关"
        },
        {
            "name": "/bids",
            "description": "标的物相关"
        },
        {
            "name": "/broadcast",
            "description": "广播相关"
        },
        {
            "name": "/chargerecord",
            "description": "充值记录相关"
        },
        {
            "name": "/config",
            "description": "配置相关"
        },
        {
            "name": "/coupon",
            "description": "优惠券相关"
        },
        {
            "name": "/dividendsub",
            "description": "标的物利息区间相关"
        },
        {
            "name": "/feedback",
            "description": "反馈相关"
        },
        {
            "name": "/finance",
            "description": "后台财务数据相关"
        },
        {
            "name": "/message",
            "description": "消息相关"
        },
        {
            "name": "/platform",
            "description": "渠道相关"
        },
        {
            "name": "/purchaserecord",
            "description": "活期购买记录相关"
        },
        {
            "name": "/rank",
            "description": "积分相关"
        },
        {
            "name": "/record",
            "description": "记录相关"
        },
        {
            "name": "/statistics",
            "description": "统计相关"
        },
        {
            "name": "/task",
            "description": "定时任务相关"
        },
        {
            "name": "/transaction",
            "description": "交易相关"
        },
        {
            "name": "/usercoupon",
            "description": "用户的优惠券相关"
        },
        {
            "name": "/users",
            "description": "用户相关"
        },
        {
            "name": "/withdrawrecord",
            "description": "提现记录相关"
        }
    ]
}
`
