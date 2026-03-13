-- a. Tampilkan data polis yang di submit setelah 15 Jan berdasarkan client yang lahir bulan september
select *
from t_policy p
join t_client c on p.client_number = c.client_number
where p.policy_submit_date > '2018-01-15' and extract(month from c.birth_date)= 9

-- b. Tampilkan data polis status INFORCE kantor di Jakarta
select *
from t_policy p
join t_agent a on p.agent_code = a.agent_code
where p.policy_status = 'INFORCE' and a.agent_office = 'JAKARTA'

-- c. Hitung column Basic Commission pada table t_agent
update t_agent a
set basic_commission = (
	select (sum(p.commission)/sum(p.premium))*100
	from t_policy p
	where p.agent_code = a.agent_code
)

select *
from t_agent

-- d. Isi kolom policy_due_date dengan kondisi tanggal akhir bulan dari 30 hari setelah policy_submit_date
update t_policy
set policy_due_date =
(date_trunc('month', policy_submit_date + 30) + interval '1 month - 1 day')::date

select * from t_policy

-- e. Tampilakn semua data produksi agent yang premiumnya setelah dipotong discount dibawah 1jt order by asc
select a.agent_code, a.agent_name, a.agent_office, p.premium, p.discount, p.premium - (p.premium * p.discount / 100) as premium_after_discount
from t_policy p
join t_agent a on p.agent_code = a.agent_code
where p.premium - (p.premium * p.discount / 100) < 1000000
order by premium_after_discount asc