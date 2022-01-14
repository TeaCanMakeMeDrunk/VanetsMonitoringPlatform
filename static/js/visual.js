//第一类动态曲线调用
dynamic_line("line1", "line2", "left");
dynamic_line("line3", "line4", "left");
dynamic_line("line5", "line6", "right");
dynamic_line("line7", "line8", "right");

//摄像头-OBU-RSU之间的线特效js
function dynamic_line(line_id1, line_id2, type) {
    //动态线的逻辑, 对应的div两个动画，底层是入队列顺序执行
    let line1 = $("#"+line_id1)
    let line2 = $("#"+line_id2)
    line1.fadeIn();
    line2.fadeIn();

    if(type==="right"){
        //右边部分线条动画
        dynamic_line_right(line1, line2, line_id1, line_id2, type);
    }else{
        //左边边部分线条动画
        dynamic_line_left(line1, line2, line_id1, line_id2, type);
    }
}

function dynamic_line_left(line1, line2, line_id1, line_id2, type) {
    let origin_left_value = line1.css("left")
    line1.animate({
        left: '+=126px',
    }, {
        duration: 1000,
        easing: 'linear',
        //上一个动画完了马上执行下一个动画(函数)
        complete: () => {
            //line1隐藏
            line1.fadeOut("normal")
            line2.animate({
                width: '80px', //斜方向利用长度来变化
            }, {
                duration: 500,
                easing: 'linear',
                complete: () => {
                    line2.fadeOut("normal")
                    //清零，下次动画开始位置
                    line1.css(type, origin_left_value)
                    line2.css("width", 0)
                    //递归调用
                    dynamic_line(line_id1, line_id2, type)
                }
            })
        }
    })
}

function dynamic_line_right(line1, line2, line_id1, line_id2, type) {
    let origin_right_value = line1.css("right")
    line1.animate({
        right: '+=125px',
    }, {
        duration: 1000,
        easing: 'linear',
        //上一个动画完了马上执行下一个动画(函数)
        complete: () => {
            //line1隐藏
            line1.fadeOut("normal")
            line2.animate({
                width: '80px', //斜方向利用长度来变化
            }, {
                duration: 500,
                easing: 'linear',
                complete: () => {
                    line2.fadeOut("normal")
                    //清零，下次动画开始位置
                    line1.css(type, origin_right_value)
                    line2.css("width", 0)
                    //递归调用
                    dynamic_line(line_id1, line_id2, type)
                }
            })
        }
    })
}

//右上角资源使用情况图表1, 获得图表需要的option
function getResRadarOption(data) {
    let color = ['#e9df3d', '#f79c19', '#21fcd6', '#08c8ff', '#df4131'];
    let max = data[0].value
    data.forEach(function (d) {
        max = d.value > max ? d.value : max;
    });

    let renderData = [{
        value: [],
        name: "系统资源使用百分比",
        symbol: 'none',
        lineStyle: {
            normal: {
                color: '#ecc03e',
                width: 2
            }
        },
        areaStyle: {
            normal: {
                color: new echarts.graphic.LinearGradient(0, 0, 1, 0,
                    [{
                        offset: 0,
                        color: 'rgba(203, 158, 24, 0.8)'
                    }, {
                        offset: 1,
                        color: 'rgba(190, 96, 20, 0.8)'
                    }], false)
            }
        }
    }];


    data.forEach(function (d, i) {
        let value = ['', '', '', '', ''];
        value[i] = max;
        renderData[0].value[i] = d.value;
        renderData.push({
            value: value,
            symbol: 'circle',
            symbolSize: 12,
            lineStyle: {
                normal: {
                    color: 'transparent'
                }
            },
            itemStyle: {
                normal: {
                    color: color[i],
                }
            }
        })
    })
    let indicator = [];

    data.forEach(function (d) {
        indicator.push({
            name: d.name,
            max: max,
            color: '#fff'
        })
    })

    return {
        tooltip: {
            show: true,
            trigger: "item"
        },
        radar: {
            center: ["50%", "50%"],//偏移位置
            radius: "80%",
            startAngle: 40, // 起始角度
            splitNumber: 4,
            shape: "circle",
            splitArea: {
                areaStyle: {
                    color: 'transparent'
                }
            },
            axisLabel: {
                show: false,
                fontSize: 20,
                color: "#000",
                fontStyle: "normal",
                fontWeight: "normal"
            },
            axisLine: {
                show: true,
                lineStyle: {
                    color: "rgba(255, 255, 255, 0.5)"
                }
            },
            splitLine: {
                show: true,
                lineStyle: {
                    color: "rgba(255, 255, 255, 0.5)"
                }
            },
            indicator: indicator
        },
        series: [{
            type: "radar",
            data: renderData
        }]
    }
}

let data = [
    {
        "name": "带宽资源",
        "value": 0
    },
    {
        "name": "上行带宽",
        "value": 0
    },
    {
        "name": "下行带宽",
        "value": 0
    },
    {
        "name": "算力资源",
        "value": 0
    },
]
resourceUsageOption = getResRadarOption(data)
////////////////////////右上角资源使用情况图表1 end