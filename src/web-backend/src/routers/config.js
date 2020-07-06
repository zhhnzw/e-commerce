export default {
    menu:[
        {key: '/', title: '概览', component: 'Home'},
        {
            key: '/goods',
            title: '商品',
            subs: [
                {key: '/goods/list', title:'商品列表', component: 'TableForGoods'},
                {key: '/order/list', title:'订单列表', component: 'TableForOrder'},
            ]
        }
    ],
    others:[],
}