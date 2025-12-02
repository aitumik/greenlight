import React from "react";

export const EventCardOld = ({
    title,
    description,
    cover,
    start_time,
    venue,
    location,
}) => {
    return (
        <div className="flex gap-3 py-2 items-center">
            <img className="h-[84px] w-[85px] object-fill rounded-md" src={cover} />
            <div>
                <h3 className="font-semibold">{title}</h3>
                <h5 className="text-sm text-neutral-500">{start_time}</h5>
                <h5 className="text-sm text-neutral-500">{venue}</h5>
            </div>
        </div>
    );
};

export const EventCard = ({
    title,
    description,
    cover,
    start_time,
    venue,
    location,
}) => {
    return (
        <div className="flex gap-3 py-2 items-center">
            <img className="h-[84px] w-[85px] object-fill rounded-md" src={cover} />
            <div>
                <h3 className="font-semibold">{title}</h3>
                <h5 className="text-sm text-neutral-500">{start_time}</h5>
                <h5 className="text-sm text-neutral-500">{venue}</h5>
            </div>
        </div>
    );
};

