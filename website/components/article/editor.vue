<template>
    <div class="golang123-editor">
        <textarea v-if="isMounted" ref="textarea"></textarea>
        <Modal v-model="modalVisible"
            title="上传图片"
            @on-ok="ok"
            @on-cancel="cancel">
            <Upload :on-success="onUploadSuccess" :name="'upFile'"
                :format="['jpg', 'jpeg', 'png', 'gif']"
                :action="uploadURL">
                <Button type="ghost" icon="ios-cloud-upload-outline">上传图片</Button>
            </Upload>
        </Modal>
    </div>
</template>

<script>
    import ErrorCode from '~/constant/ErrorCode'
    import config from '~/config'

    export default {
        props: {
            value: {
                type: String,
                required: true
            }
        },
        data () {
            return {
                isMounted: false,
                host: '',
                simplemde: null,
                SimpleMDE: null,
                uploadURL: config.apiURL + '/upload',
                modalVisible: false,
                toolbar: null
            }
        },
        methods: {
            ok () {
                this.$Message.info('点击了确定')
            },
            cancel () {
                this.$Message.info('点击了取消')
            },
            showUpload () {
                this.modalVisible = true
            },
            onUploadSuccess (res, file) {
                if (res) {
                    if (res.errNo === ErrorCode.SUCCESS) {
                        var url = 'https://' + this.host + res.data.url
                        this.simplemde.setImageURL(url)
                        this.SimpleMDE.drawImage(this.simplemde)
                    } else if (res.errNo === ErrorCode.LOGIN_TIMEOUT) {
                        location.href = '/signin'
                    }
                }
            }
        },
        head () {
            return {
                script: [
                    { src: '/javascripts/editor/simplemde.js' }
                ]
            }
        },
        mounted () {
            this.isMounted = true
            this.$nextTick(function () {
                this.host = document.location.hostname
                let SimpleMDE = window.SimpleMDE
                this.toolbar = [
                    {
                        name: 'bold',
                        action: SimpleMDE.toggleBold,
                        className: 'fa fa-bold',
                        title: '粗体'
                    },
                    {
                        name: 'italic',
                        action: SimpleMDE.toggleItalic,
                        className: 'fa fa-italic',
                        title: '斜体'
                    },
                    {
                        name: 'heading',
                        action: SimpleMDE.toggleHeadingSmaller,
                        className: 'fa fa-header',
                        title: '标题'
                    },
                    '|',
                    {
                        name: 'unordered-list',
                        action: SimpleMDE.toggleUnorderedList,
                        className: 'fa fa-list-ul',
                        title: '无序列表'
                    },
                    {
                        name: 'ordered-list',
                        action: SimpleMDE.toggleOrderedList,
                        className: 'fa fa-list-ol',
                        title: '有序列表'
                    },
                    {
                        name: 'table',
                        action: SimpleMDE.drawTable,
                        className: 'fa fa-table',
                        title: '表格'
                    },
                    {
                        name: 'horizontal-rule',
                        action: SimpleMDE.drawHorizontalRule,
                        className: 'fa fa-minus',
                        title: '水平分隔线'
                    },
                    '|',
                    {
                        name: 'code',
                        action: SimpleMDE.toggleCodeBlock,
                        className: 'fa fa-code',
                        title: '代码'
                    },
                    {
                        name: 'quote',
                        action: SimpleMDE.toggleBlockquote,
                        className: 'fa fa-quote-left',
                        title: '块引用'
                    },
                    {
                        name: 'link',
                        action: SimpleMDE.drawLink,
                        className: 'fa fa-link',
                        title: '链接'
                    },
                    {
                        name: 'image',
                        action: this.showUpload,
                        className: 'fa fa-picture-o',
                        title: '上传图片'
                    },
                    '|',
                    {
                        name: 'preview',
                        action: SimpleMDE.togglePreview,
                        className: 'fa fa-eye no-disable',
                        title: '预览'
                    },
                    {
                        name: 'fullscreen',
                        action: SimpleMDE.toggleFullScreen,
                        className: 'fa fa-arrows-alt no-disable no-mobile',
                        title: '全屏'
                    }
                ]
                this.SimpleMDE = SimpleMDE
                var simplemde = new SimpleMDE({
                    element: this.$refs.textarea,
                    promptURLs: false,
                    spellChecker: false,
                    toolbar: this.toolbar
                })
                this.simplemde = simplemde
                var pt = SimpleMDE.prototype
                if (!pt.getImageURL) {
                    pt.getImageURL = function () {
                        return this.imageURL
                    }
                    pt.setImageURL = function (url) {
                        this.imageURL = url
                    }
                }
                var self = this
                this.simplemde.codemirror.on('change', function () {
                    self.$emit('change', self.simplemde.value())
                })
                this.simplemde.value(this.value)
            })
        }
    }
</script>
