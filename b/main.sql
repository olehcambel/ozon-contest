select g.id, g.name
from goods g
    inner join (select count(*) as len from tags) as mx
where mx.len = (
    select count(*)
    from tags_goods tg
    where tg.goods_id = g.id
)