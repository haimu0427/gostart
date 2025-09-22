(function () {
    // ================= 基础请求封装 =================
    var hasAxios = !!window.axios;
    var axiosInstance = hasAxios ? window.axios.create({}) : null;
    if (!hasAxios) {
        console.warn('[login] 未找到全局 axios, 使用 fetch 降级实现');
    } else {
        console.log('[login] axios 已加载, 创建独立实例用于登录与注册');
    }

    function normalizeResp(res) {
        // axios: res.data 是真正数据； fetch: res已是对象
        var data = (res && res.data) ? res.data : res;
        return data || {};
    }

    function postJSON(url, data) {
        if (axiosInstance) {
            return axiosInstance.post(url, data).then(normalizeResp);
        }
        return fetch(url, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(data)
        }).then(function (r) { return r.json(); }).then(normalizeResp);
    }

    function setGlobalAuthHeader(token) {
        if (token && window.axios) {
            // 设置全局默认头 + 拦截器确保后续请求自动携带
            window.axios.defaults.headers.common['Authorization'] = 'Bearer ' + token;
        }
    }
    new Vue({
        el: '#app',
        data: function () {
            var validateConfirm = (rule, value, callback) => {
                if (value !== this.register.password) {
                    callback(new Error('两次输入的密码不一致'));
                } else { callback(); }
            };
            return {
                loading: false,
                error: '',
                form: { username: '', password: '' },
                rules: {
                    username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
                    password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
                },
                showRegister: false,
                registerLoading: false,
                register: { username: '', password: '', confirm: '' },
                registerRules: {
                    username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
                    password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
                    confirm: [{ required: true, message: '请再次输入密码', trigger: 'blur' }, { validator: validateConfirm, trigger: 'blur' }]
                }
            }
        },
        methods: {
            onSubmit: function () {
                var self = this;
                if (self.loading) return; // 防重复
                this.$refs.loginForm.validate(function (valid) {
                    if (!valid) return;
                    if (self.form.password.length < 3) {
                        self.error = '密码长度至少3位';
                        return;
                    }
                    self.loading = true; self.error = '';
                    console.log('[login] submit login', self.form.username);
                    postJSON('/login', self.form)
                        .then(function (data) {
                            if (data.code === 200) {
                                if (data.token) {
                                    localStorage.setItem('auth_token', data.token);
                                    setGlobalAuthHeader(data.token);
                                }
                                // 延迟 100ms 给本地 header 设置时间（可选）
                                setTimeout(function () { window.location.href = '/'; }, 50);
                            } else {
                                self.error = data.message || '登录失败';
                            }
                        })
                        .catch(function (err) {
                            console.error('[login] login error', err);
                            self.error = (err && err.message) || '请求错误';
                        })
                        .finally(function () { self.loading = false; });
                });
            },
            onRegister: function () {
                var self = this;
                this.$refs.registerForm.validate(function (valid) {
                    if (!valid) return;
                    self.registerLoading = true;
                    console.log('[login] submit register', self.register.username);
                    if (self.register.password.length < 3) {
                        self.$message.error('密码长度至少3位');
                        self.registerLoading = false;
                        return;
                    }
                    postJSON('/register', { username: self.register.username, password: self.register.password })
                        .then(function (data) {
                            if (data.code === 200) {
                                self.$message.success('注册成功, 请登录');
                                self.showRegister = false;
                                self.form.username = self.register.username;
                                self.form.password = '';
                                self.resetRegister();
                            } else {
                                self.$message.error(data.message || '注册失败');
                            }
                        })
                        .catch(function (err) {
                            console.error('[login] register error', err);
                            self.$message.error((err && err.message) || '请求错误');
                        })
                        .finally(function () { self.registerLoading = false; });
                })
            },
            resetRegister: function () {
                this.register = { username: '', password: '', confirm: '' };
                if (this.$refs.registerForm) { this.$refs.registerForm.clearValidate(); }
            }
        },
        mounted: function () {
            var tk = localStorage.getItem('auth_token');
            if (tk) {
                setGlobalAuthHeader(tk);
                // 已登录, 直接去主页（可以注释掉允许在登录页切换账号）
                window.location.href = '/';
            }
        }
    });
})();