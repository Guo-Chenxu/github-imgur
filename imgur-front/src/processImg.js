(function (window) {

    /**
     * 
     * 作者：混沌传奇
     * 
     * 邮箱地址：iot-pro_lizeng@foxmail.com
     * 
     * 日期：2017-10-26
     * 
     * 插件功能：压缩图片&&纠正图片方向&&返回二进制(Blob)图片元数据组成的列表
     * 
     */

    window.lzImgProcess = function () {


        var Orientation = '', //图片方向角
            base64List = [], //压缩后的二进制图片数据列表

            canvas = document.createElement("canvas"), //用于压缩图片（纠正图片方向）的canvas
            ctx = canvas.getContext('2d'),

            file_type = 'image/jpeg', //图片类型
            qlty = 0.5, //图片压缩品质，默认是0.5，可选范围是0-1的数字类型的值，可配置
            imgWH = 1000; //压缩后的图片的最大宽度和高度，默认是1000px，可配置


        /**
         * @actionName process, 
         *        方法功能：压缩图片&&纠正图片方向&&返回二进制(Blob)图片元数据
         * 
         * @param fileList,传入函数的文件列表对象，fileList对象是来自用户在一个<input>元素上选择文件后返回的FileList对象
         *        注意：图片类型必须是jpeg||png
         *        比如：<input id="uploadImage" onchange="loadImageFile();" /> 
         *              function loadImageFile() {
         *                //获取返回的fileList对象
         *                var fileList = document.getElementById("uploadImage").files;
         *              }
         * @param getBlobList [Blob],获取压缩结果的钩子函数，接受一个参数。
         *        功能：在图片压缩完毕后，获取压缩后的二进制图片数据对象组成的数组，参数即：压缩后的二进制图片数据(blob)组成的list
         *        
         * @param quality,传入函数的图片压缩比率(品质)，可选范围0-1的数字类型的值，默认是0.5
         *
         * @param WH,传入函数的图片压缩后的最大图片宽度和高度，默认是1000，单位是px，可自由配置。
         *        注意：最好不要超过1000，数字过大，容易导致canvas压缩失败。由于没做瓦片处理，所以有这个限制。1000*1000的图片在前端中，基本也够用了。
         *        
         */
        function process(fileList, getBase64List, quality, WH) {
            base64List = []; //初始化base64List
            // 判断参数fileList的长度是否大于0
            if (!fileList.length) {
                console.log('警告：传进方法process的参数fileList长度必须大于零！！！')
                return;
            }

            //如果quality参数有值，则把quality赋值给qlty(图片压缩的品质)
            if (quality)
                qlty = quality;

            //如果WH参数有值，则把WH赋值给imgWH（压缩后的图片的最大宽度和高度）
            if (WH && WH < 1000 && WH > 0) {
                imgWH = WH;
            }

            // 把传进来的fileList转为数组类型
            var files = Array.prototype.slice.call(fileList);

            files.forEach(function (file, i) {
                if (!/\/(?:jpeg|png)/i.test(file.type)) {
                    console.log('警告：图片必须是jpeg||png类型！！！');
                    return;
                }

                var reader = new FileReader();

                reader.onload = function () {
                    var img = new Image();
                    img.src = this.result;

                    // 图片加载完毕之后进行压缩
                    if (img.complete) {
                        callback();
                    } else {
                        img.onload = callback;
                    }

                    function callback() {
                        //获取照片方向角属性，用户旋转控制  
                        EXIF.getData(img, function () {
                            // alert(EXIF.pretty(this));
                            EXIF.getAllTags(this);
                            // alert(EXIF.getTag(this, 'Orientation'));
                            Orientation = EXIF.getTag(this, 'Orientation');
                            console.log('Orientation:', Orientation)
                            if (Orientation == "" || Orientation == undefined || Orientation == null) {
                                Orientation = 1;
                            }
                        });

                        //将图片存入列表中
                        base64List.push(img.src.split(',')[1]);

                        //将图片数据对象组成的list通过钩子函数返回出去
                        if (base64List.length === files.length) {
                            if (getBase64List)
                                getBase64List(base64List);
                        }

                        img = null;
                    }

                };

                reader.readAsDataURL(file);
            })
        }

        /**
         * 返回一个process方法
         * 
         * process方法：处理图片数据
         * 
         */
        return process;

    }
})(window)

