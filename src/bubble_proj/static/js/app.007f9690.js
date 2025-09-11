(function (modules) {
    // webpackBootstrap
    var installedModules = {};
    var installedChunks = { app: 0 };
    var deferredModules = [];

    function webpackJsonpCallback(data) {
        var chunkIds = data[0];
        var moreModules = data[1];
        var executeModules = data[2];
        var i = 0, resolves = [];
        for (; i < chunkIds.length; i++) {
            var chunkId = chunkIds[i];
            if (Object.prototype.hasOwnProperty.call(installedChunks, chunkId) && installedChunks[chunkId]) {
                resolves.push(installedChunks[chunkId][0]);
            }
            installedChunks[chunkId] = 0;
        }
        for (var moduleId in moreModules) {
            if (Object.prototype.hasOwnProperty.call(moreModules, moduleId)) {
                modules[moduleId] = moreModules[moduleId];
            }
        }
        if (parentJsonpFunction) parentJsonpFunction(data);
        while (resolves.length) {
            resolves.shift()();
        }
        deferredModules.push.apply(deferredModules, executeModules || []);
        return checkDeferredModules();
    }

    function checkDeferredModules() {
        var result;
        for (var i = 0; i < deferredModules.length; i++) {
            var deferredModule = deferredModules[i];
            var fulfilled = true;
            for (var j = 1; j < deferredModule.length; j++) {
                var depId = deferredModule[j];
                if (installedChunks[depId] !== 0) fulfilled = false;
            }
            if (fulfilled) {
                deferredModules.splice(i--, 1);
                result = __webpack_require__(__webpack_require__.s = deferredModule[0]);
            }
        }
        return result;
    }

    function __webpack_require__(moduleId) {
        if (installedModules[moduleId]) return installedModules[moduleId].exports;
        var module = installedModules[moduleId] = { i: moduleId, l: false, exports: {} };
        modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);
        module.l = true;
        return module.exports;
    }

    __webpack_require__.m = modules;
    __webpack_require__.c = installedModules;
    __webpack_require__.d = function (exports, name, getter) {
        if (!__webpack_require__.o(exports, name)) {
            Object.defineProperty(exports, name, { enumerable: true, get: getter });
        }
    };
    __webpack_require__.r = function (exports) {
        if (typeof Symbol !== "undefined" && Symbol.toStringTag) {
            Object.defineProperty(exports, Symbol.toStringTag, { value: "Module" });
        }
        Object.defineProperty(exports, "__esModule", { value: true });
    };
    __webpack_require__.t = function (value, mode) {
        if (mode & 1) value = __webpack_require__(value);
        if (mode & 8) return value;
        if (mode & 4 && typeof value === "object" && value && value.__esModule) return value;
        var ns = Object.create(null);
        __webpack_require__.r(ns);
        Object.defineProperty(ns, "default", { enumerable: true, value: value });
        if (mode & 2 && typeof value != "string")
            for (var key in value) __webpack_require__.d(ns, key, function (key) { return value[key]; }.bind(null, key));
        return ns;
    };
    __webpack_require__.n = function (module) {
        var getter = module && module.__esModule ?
            function getDefault() { return module['default']; } :
            function getModuleExports() { return module; };
        __webpack_require__.d(getter, 'a', getter);
        return getter;
    };
    __webpack_require__.o = function (object, property) {
        return Object.prototype.hasOwnProperty.call(object, property);
    };
    __webpack_require__.p = "/";

    var jsonpArray = window["webpackJsonp"] = window["webpackJsonp"] || [];
    var oldJsonpFunction = jsonpArray.push.bind(jsonpArray);
    jsonpArray.push = webpackJsonpCallback;
    jsonpArray = jsonpArray.slice();
    for (var i = 0; i < jsonpArray.length; i++) webpackJsonpCallback(jsonpArray[i]);
    var parentJsonpFunction = oldJsonpFunction;
    deferredModules.push([0, "chunk-vendors"]);
    checkDeferredModules();
})({
    0: function (module, exports, __webpack_require__) {
        module.exports = __webpack_require__("56d7");
    },
    "034f": function (module, exports, __webpack_require__) {
        "use strict";
        var style = __webpack_require__("85ec"), styleModule = __webpack_require__.n(style);
        styleModule.a;
    },
    "56d7": function (module, exports, __webpack_require__) {
        "use strict";
        __webpack_require__.r(exports);
        __webpack_require__("e260");
        __webpack_require__("e6cf");
        __webpack_require__("cca6");
        __webpack_require__("a79d");
        var Vue = __webpack_require__("2b0e"),
            axiosLib = (__webpack_require__("d3b7"), __webpack_require__("bc3a")),
            axios = __webpack_require__.n(axiosLib),
            axiosConfig = {},
            axiosInstance = axios.a.create(axiosConfig);

        axiosInstance.interceptors.request.use(
            function (config) { return config; },
            function (error) { return Promise.reject(error); }
        );
        axiosInstance.interceptors.response.use(
            function (response) { return response; },
            function (error) { return Promise.reject(error); }
        );

        var Plugin = {};
        Plugin.install = function (Vue) {
            Vue.axios = axiosInstance;
            window.axios = axiosInstance;
            Object.defineProperties(Vue.prototype, {
                axios: { get: function () { return axiosInstance; } },
                $axios: { get: function () { return axiosInstance; } }
            });
        };
        Vue["default"].use(Plugin);

        var AppTemplate = function () {
            var vm = this, createElement = vm.$createElement, c = vm._self._c || createElement;
            return c("Index");
        }, appStaticRenderFns = [];

        var IndexTemplate = function () {
            var vm = this, createElement = vm.$createElement, c = vm._self._c || createElement;
            return c("el-container", [
                c("el-header", [vm._v("gin框架小练习")]),
                c("el-main", [
                    c("el-row", { attrs: { type: "flex", justify: "center" } }, [
                        c("el-col", { attrs: { xs: 20, span: 12 } }, [
                            c("div", { staticClass: "grid-content" }, [
                                c("el-divider", [c("h1", [vm._v("bubble清单")])]),
                                c("TodoList")
                            ], 1)
                        ])
                    ], 1)
                ], 1),
                c("el-footer", [vm._v("q1mi出品 Go学习交流QQ群：645090316")])
            ], 1);
        }, indexStaticRenderFns = [];

        var TodoListTemplate = function () {
            var vm = this, createElement = vm.$createElement, c = vm._self._c || createElement;
            return c("el-card", { staticClass: "box-card" }, [
                c("el-row", { attrs: { gutter: 20 } }, [
                    c("el-col", { attrs: { span: 16, offset: 2 } }, [
                        c("el-input", {
                            attrs: { size: "", placeholder: "请输入待办事项..." },
                            model: {
                                value: vm.newTitle,
                                callback: function (val) { vm.newTitle = val; },
                                expression: "newTitle"
                            }
                        })
                    ], 1),
                    c("el-col", { attrs: { span: 6 } }, [
                        c("el-button", {
                            attrs: { type: "primary", icon: "el-icon-plus", circle: "" },
                            on: { click: vm.handleAdd }
                        })
                    ], 1)
                ], 1),
                c("el-divider"),
                c("el-table", {
                    staticStyle: { width: "100%" },
                    attrs: { data: vm.tableData, "row-class-name": vm.tableRowClassName }
                }, [
                    c("el-table-column", { attrs: { type: "index", width: "50" } }),
                    c("el-table-column", { attrs: { align: "center", label: "待办事项", prop: "title" } }),
                    c("el-table-column", {
                        attrs: { align: "right", label: "操作" },
                        scopedSlots: vm._u([{
                            key: "default",
                            fn: function (scope) {
                                return [
                                    c("el-button", {
                                        directives: [{ name: "show", rawName: "v-show", value: !scope.row.status, expression: "!scope.row.status" }],
                                        attrs: { type: "success", icon: "el-icon-check", circle: "" },
                                        on: { click: function () { return vm.handleEdit(scope.$index, scope.row); } }
                                    }),
                                    c("el-button", {
                                        directives: [{ name: "show", rawName: "v-show", value: scope.row.status, expression: "scope.row.status" }],
                                        attrs: { type: "warning", icon: "el-icon-refresh-left", circle: "" },
                                        on: { click: function () { return vm.handleEdit(scope.$index, scope.row); } }
                                    }),
                                    c("el-button", {
                                        attrs: { type: "danger", icon: "el-icon-close", circle: "" },
                                        on: { click: function () { return vm.handleDelete(scope.$index, scope.row.id); } }
                                    })
                                ];
                            }
                        }])
                    })
                ], 1)
            ], 1);
        }, todoListStaticRenderFns = [];

        var TodoList = {
            name: "TodoList",
            data: function () {
                return {
                    tableData: [],
                    newTitle: ""
                };
            },
            mounted: function () {
                var vm = this;
                this.axios.get("/v1/todo").then(function (res) {
                    vm.tableData = res.data;
                });
            },
            methods: {
                tableRowClassName: function (rowObj) {
                    var row = rowObj.row;
                    return row.status ? "success-row" : "";
                },
                getTodoList: function () {
                    var vm = this;
                    this.axios.get("/v1/todo").then(function (res) {
                        vm.tableData = res.data;
                    });
                },
                handleEdit: function (index, row) {
                    var vm = this,
                        msg = row.status ? " 置为未完成" : " 置为已完成";
                    this.axios.put("/v1/todo/" + row.id, { status: !row.status }).then(function () {
                        vm.tableData[index].status = !row.status;
                        vm.$message({
                            showClose: true,
                            duration: 1500,
                            message: "<" + row.title + "> " + msg,
                            type: "success"
                        });
                    });
                },
                handleDelete: function (index, id) {
                    var vm = this;
                    this.axios.delete("/v1/todo/" + id).then(function () {
                        vm.tableData.splice(index, 1);
                        vm.$message({
                            showClose: true,
                            duration: 1500,
                            message: "删除待办事项成功",
                            type: "success"
                        });
                    });
                },
                handleAdd: function () {
                    var vm = this;
                    if (this.newTitle != "") {
                        this.axios.post("/v1/todo", { title: this.newTitle }).then(function () {
                            vm.getTodoList();
                            vm.$message({
                                showClose: true,
                                duration: 1500,
                                message: "添加待办事项成功",
                                type: "success"
                            });
                        });
                        this.newTitle = "";
                    } else {
                        this.$message({
                            showClose: true,
                            duration: 1500,
                            message: "title不能为空哦",
                            type: "warning"
                        });
                    }
                }
            }
        };

        var TodoListComponent = TodoList,
            TodoListRender = (__webpack_require__("ed30"), __webpack_require__("2877")),
            TodoListExport = Object(TodoListRender["a"])(TodoListComponent, TodoListTemplate, todoListStaticRenderFns, false, null, null, null),
            TodoListFinal = TodoListExport.exports;

        var Index = {
            name: "Index",
            components: { TodoList: TodoListFinal }
        };

        var IndexComponent = Index,
            IndexExport = (__webpack_require__("8fc1"), Object(TodoListRender["a"])(IndexComponent, IndexTemplate, indexStaticRenderFns, false, null, null, null)),
            IndexFinal = IndexExport.exports;

        var App = {
            name: "app",
            components: { Index: IndexFinal }
        };

        var AppComponent = App,
            AppExport = (__webpack_require__("034f"), Object(TodoListRender["a"])(AppComponent, AppTemplate, appStaticRenderFns, false, null, null, null)),
            AppFinal = AppExport.exports;

        var ElementUI = __webpack_require__("5c96"),
            ElementUIModule = __webpack_require__.n(ElementUI);
        __webpack_require__("0fae");
        Vue["default"].use(ElementUIModule.a);

        var VueRouter = __webpack_require__("8c4f");
        Vue["default"].use(VueRouter["a"]);
        var routes = [{ path: "/", name: "index", component: IndexFinal }],
            router = new VueRouter["a"]({ routes: routes });

        Vue["default"].config.productionTip = true;
        new Vue["default"]({
            router: router,
            render: function (h) { return h(AppFinal); }
        }).$mount("#app");
    },
    "85ec": function (module, exports, __webpack_require__) { },
    "89d2": function (module, exports, __webpack_require__) { },
    "8fc1": function (module, exports, __webpack_require__) {
        "use strict";
        var style = __webpack_require__("9272"), styleModule = __webpack_require__.n(style);
        styleModule.a;
    },
    9272: function (module, exports, __webpack_require__) { },
    ed30: function (module, exports, __webpack_require__) {
        "use strict";
        var style = __webpack_require__("89d2"), styleModule = __webpack_require__.n(style);
        styleModule.a;
    }
});