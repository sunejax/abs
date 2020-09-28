function allocate(rec, field, inventory)
    if (not (aerospike:exists(rec))) then
        rc = aerospike:create(rec)
        if (rc == nil or rec == nil) then
            warn("[ERROR]<%s> Problems creating record", meth);
            error("ERROR creating record");
        end
    end

    local count = 1;
    local ret = map();
    ret['overFlow'] = false;

    --check if the rec exists
    if (rec[field] == nil) then
        count = 1; -- this is the first one
    else
        --if it does then increment
        count = rec[field];
        count = count + 1;
    end
    --Updates Record
    rec[field] = count;

    if (count > inventory) then
        rec['overFlow'] = 1;
        ret['overFlow'] = true;
        return ret;
    end

    ret['success'] = aerospike:update(rec);

    --Returns
    return ret;
end

